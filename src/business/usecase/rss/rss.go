package rss

import (
	"context"

	"github.com/linggaaskaedo/go-play/src/business/domain/rss"
	"github.com/xtfly/log4g/api"
)

type UsecaseItf interface {
	GetLatestNews(ctx context.Context) error
}

type rssUsecase struct {
	logger api.Logger
	opt    Options
	rss    rss.DomainItf
}

type Options struct {
	RSSUrl string
}

func InitRSSUsecase(
	logger api.Logger,
	opt Options,
	rss rss.DomainItf,
) UsecaseItf {
	return &rssUsecase{
		logger: logger,
		opt:    opt,
		rss:    rss,
	}
}
