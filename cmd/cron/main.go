package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/henriquepw/prata-api/internal/database"
	"github.com/henriquepw/prata-api/internal/domains/recurrence"
	"github.com/henriquepw/prata-api/internal/domains/transaction"
	"github.com/henriquepw/prata-api/internal/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

const (
	CronEveryday  = "0 0 * * *"
	CronEveryHour = "0 * * * *"
)

type jobServer struct {
	ctx             context.Context
	scheduler       gocron.Scheduler
	recurrenceStore recurrence.RecurrenceStore
	transactionSVC  transaction.TransactionService
}

func New(ctx context.Context, db *sqlx.DB) (*jobServer, error) {
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
		return nil, err
	}

	recurrenceStore := recurrence.NewStore(db)
	transactionStore := transaction.NewStore(db)
	transactionSVC := transaction.NewService(transactionStore)

	return &jobServer{
		scheduler:       s,
		recurrenceStore: recurrenceStore,
		transactionSVC:  transactionSVC,
	}, nil
}

func (s *jobServer) addTask(cron, name string, task func() error) error {
	_, err := s.scheduler.NewJob(
		gocron.CronJob(cron, false),
		gocron.NewTask(task),
		gocron.WithName(name),
	)
	if err != nil {
		log.Error("can't setup job", "name", name)
	}

	return err
}

func (s *jobServer) Start() error {
	err := s.addTask(CronEveryday, "create-transactions-by-transactions", s.createTodayTransactions)
	if err != nil {
		return err
	}

	s.scheduler.Start()
	<-s.ctx.Done()

	return s.scheduler.Shutdown()
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	db, err := database.GetDB()
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		return err
	}
	defer db.Close()

	if os.Getenv(env.Version) == "DEVELOP" {
		db.SetMaxOpenConns(1)
	}

	jobServer, err := New(ctx, db)
	if err != nil {
		return err
	}

	return jobServer.Start()
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("Cronjob finished with error", "error", err.Error())
		os.Exit(1)
	}
}
