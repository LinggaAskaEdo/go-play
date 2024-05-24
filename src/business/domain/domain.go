package domain

import (
	"database/sql"

	"github.com/linggaaskaedo/go-play/src/business/domain/rss"
	"github.com/xtfly/log4g/api"
)

type Domain struct {
	RSS rss.DomainItf
}

func Init(logger api.Logger, sqlClient0 *sql.DB) *Domain {
	return &Domain{
		RSS: rss.InitRSSDomain(logger, sqlClient0),
	}
}
