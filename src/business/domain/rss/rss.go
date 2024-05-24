package rss

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/linggaaskaedo/go-play/src/business/entity"
	"github.com/xtfly/log4g/api"
)

type DomainItf interface {
	GetNewsByUrl(ctx context.Context, url string) (bool, error)
	CreateNews(ctx context.Context, v entity.NewsArticle) (entity.NewsArticle, error)
}

type rssDomain struct {
	logger api.Logger
	sql0   *sqlx.DB
}

func InitRSSDomain(
	logger api.Logger,
	sql0 *sqlx.DB,
) DomainItf {
	return &rssDomain{
		logger: logger,
		sql0:   sql0,
	}
}
