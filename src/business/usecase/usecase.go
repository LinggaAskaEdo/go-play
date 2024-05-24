package usecase

import (
	"github.com/linggaaskaedo/go-play/src/business/domain"
	"github.com/linggaaskaedo/go-play/src/business/usecase/rss"
	"github.com/xtfly/log4g/api"
)

type Usecase struct {
	RSS rss.UsecaseItf
}

type Options struct {
	RSS rss.Options
}

func Init(
	logger api.Logger,
	options Options,
	dom *domain.Domain,
) *Usecase {
	return &Usecase{
		RSS: rss.InitRSSUsecase(
			logger,
			options.RSS,
			dom.RSS,
		),
	}
}
