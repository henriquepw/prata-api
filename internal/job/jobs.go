package job

import (
	"log"
	"log/slog"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

const (
	CronEveryday  = "0 0 * * *"
	CronEveryHour = "0 * * * *"
)

type jobServer struct {
	scheduler gocron.Scheduler
	running   map[string]struct{}
}

func New() *jobServer {
	s, err := gocron.NewScheduler(
		gocron.WithStopTimeout(3*time.Hour),
		gocron.WithGlobalJobOptions(
			gocron.WithEventListeners(
				gocron.BeforeJobRuns(func(jobID uuid.UUID, jobName string) {
					slog.Info("[JOB] Starting:", "jobID", jobID, "jobName", jobName)
				}),
				gocron.AfterJobRuns(func(jobID uuid.UUID, jobName string) {
					slog.Info("[JOB] Finished:", "jobID", jobID, "jobName", jobName)
				}),
				gocron.AfterJobRunsWithError(func(jobID uuid.UUID, jobName string, err error) {
					slog.Info("[JOB] Finished with error:", "jobID", jobID, "jobName", jobName, "error", err.Error())
				}),
			),
		))
	if err != nil {
		panic(err)
	}

	return &jobServer{
		scheduler: s,
		running:   map[string]struct{}{},
	}
}

func (s *jobServer) addTask(cron, name string, task func()) {
	_, err := s.scheduler.NewJob(gocron.CronJob(cron, false), gocron.NewTask(func() {
		_, ok := s.running[name]
		if ok {
			return
		}

		s.running[name] = struct{}{}
		task()
		delete(s.running, name)
	}), gocron.WithName(name))
	if err != nil {
		slog.Error("can't start job", "name", name)
		panic(err)
	}
}

func (s *jobServer) Start() error {
	log.Print("TODO JOB")

	s.addTask(CronEveryday, "create-transactions-by-transactions", createTransactionByRecurrence)

	return nil
}
