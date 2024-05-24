package rss

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-play/src/business/entity"
	"github.com/xtfly/log4g/api"
)

type DomainItf interface {
	GetNewsByUrl(ctx context.Context, url string) (bool, error)
	CreateNews(ctx context.Context, v entity.NewsArticle) (entity.NewsArticle, error)
}

type rssDomain struct {
	logger api.Logger
	sql0   *sql.DB
}

func InitRSSDomain(logger api.Logger, sql0 *sql.DB) DomainItf {
	return &rssDomain{logger: logger, sql0: sql0}
}
