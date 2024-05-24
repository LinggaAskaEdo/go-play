package rss

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-play/src/business/entity"
)

func (r *rssDomain) GetNewsByUrl(ctx context.Context, url string) (bool, error) {
	result, err := r.getNewsByUrl(ctx, url)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *rssDomain) CreateNews(ctx context.Context, v entity.NewsArticle) (entity.NewsArticle, error) {
	// Get a Tx for making transaction requests
	tx, err := r.sql0.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	// tx, err := r.sql0.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error(err)

		return v, err
	}
	defer tx.Rollback()

	// Create News
	tx, v, err = r.createSQLNews(ctx, tx, v)
	if err != nil {
		r.logger.Error(err)
		tx.Rollback()

		return v, err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return v, err
	}

	return v, nil
}
