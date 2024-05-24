package scheduler

import (
	"context"
	"sync"

	"github.com/go-co-op/gocron/v2"
	"github.com/linggaaskaedo/go-play/src/business/usecase"
	"github.com/xtfly/log4g/api"
)

var once = &sync.Once{}

type Scheduler interface{}

type Options struct {
	RSS RSSOptions
}

type RSSOptions struct {
	Enabled bool
	Period  string
}

type scheduler struct {
	logger api.Logger
	opt    Options
	uc     *usecase.Usecase
}

func Init(
	logger api.Logger,
	opt Options,
	uc *usecase.Usecase,
) Scheduler {
	var e Scheduler

	once.Do(func() {
		e := &scheduler{
			logger: logger,
			opt:    opt,
			uc:     uc,
		}

		e.Serve()
	})

	return e
}

func (s *scheduler) Serve() {
	ctx := context.Background()

	if s.opt.RSS.Enabled {
		rssScheduler, err := gocron.NewScheduler()
		if err != nil {
			s.logger.Error(err)
		}

		// add a job to the scheduler
		rssJob, err := rssScheduler.NewJob(
			gocron.CronJob(
				s.opt.RSS.Period,
				false,
			),
			gocron.NewTask(
				func(ctx context.Context) {
					s.uc.RSS.GetLatestNews(ctx)
				},
				ctx,
			),
		)
		if err != nil {
			s.logger.Error(err)
		}

		rssScheduler.Start()
		err = rssJob.RunNow()
		if err != nil {
			s.logger.Error(err)
		}
	}
}
