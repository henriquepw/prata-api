package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/henriquepw/prata-api/internal/plataform/database"
	"github.com/henriquepw/prata-api/internal/plataform/env"
	"github.com/henriquepw/prata-api/internal/recurrence"
	"github.com/henriquepw/prata-api/internal/transaction"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

const (
	CronEveryday  = "0 0 * * *"
	CronEveryHour = "0 * * * *"
)

type jobServer struct {
	scheduler       gocron.Scheduler
	recurrenceStore recurrence.RecurrenceStore
	transactionSVC  transaction.TransactionService
}

func New(db *sqlx.DB) *jobServer {
	s, err := gocron.NewScheduler(
		gocron.WithStopTimeout(3*time.Hour),
		gocron.WithGlobalJobOptions(
			gocron.WithSingletonMode(gocron.LimitModeReschedule),
			gocron.WithEventListeners(
				gocron.BeforeJobRuns(func(jobID uuid.UUID, jobName string) {
					log.Info("[JOB] Starting:", "jobID", jobID, "jobName", jobName)
				}),
				gocron.AfterJobRuns(func(jobID uuid.UUID, jobName string) {
					log.Info("[JOB] Finished:", "jobID", jobID, "jobName", jobName)
				}),
				gocron.AfterJobRunsWithError(func(jobID uuid.UUID, jobName string, err error) {
					log.Info("[JOB] Finished with error:", "jobID", jobID, "jobName", jobName, "error", err.Error())
				}),
			),
		))
	if err != nil {
		panic(err)
	}

	recurrenceStore := recurrence.NewStore(db)
	transactionStore := transaction.NewStore(db)
	transactionSVC := transaction.NewService(transactionStore)

	return &jobServer{
		scheduler:       s,
		recurrenceStore: recurrenceStore,
		transactionSVC:  transactionSVC,
	}
}

func (s *jobServer) addTask(cron, name string, task func() error) {
	_, err := s.scheduler.NewJob(
		gocron.CronJob(cron, false),
		gocron.NewTask(task),
		gocron.WithName(name),
	)
	if err != nil {
		log.Error("can't setup job", "name", name)
		panic(err)
	}
}

func (s *jobServer) Start() error {
	s.addTask(CronEveryday, "create-transactions-by-transactions", s.createTransactionByRecurrence)

	return nil
}

func main() {
	db, err := database.GetDB()
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		return
	}
	defer db.Close()

	if os.Getenv(env.Version) == "DEVELOP" {
		db.SetMaxOpenConns(1)
	}

	jobServer := New(db)
	if err := jobServer.Start(); err != nil {
		slog.Error("failed to start job server", "error", err)
		return
	}
}
