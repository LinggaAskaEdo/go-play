package rss

import (
	"context"
	"fmt"

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
	// Create a helper function for preparing failure results
	fail := func(err error) (entity.NewsArticle, error) {
		return v, fmt.Errorf("CreateNews: %v", err)
	}

	// Get a Tx for making transaction requests
	// tx, err := r.sql.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	tx, err := r.sql0.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	// Create News
	tx, v, err = r.createSQLNews(ctx, tx, v)
	if err != nil {
		tx.Rollback()
		return fail(err)
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return fail(err)
	}

	return v, nil
}
