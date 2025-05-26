package job

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/henriquepw/prata-api/internal/domains/recurrence"
	"github.com/henriquepw/prata-api/internal/domains/transaction"
	"github.com/jmoiron/sqlx"
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
