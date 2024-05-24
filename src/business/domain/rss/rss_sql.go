package rss

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/linggaaskaedo/go-play/src/business/entity"
)

func (r *rssDomain) getNewsByUrl(ctx context.Context, url string) (bool, error) {
	var isExist bool

	err := r.sql0.QueryRowContext(ctx, GetNewsByUrl, url).Scan(&isExist)
	if err != nil {
		r.logger.Error("Error when ", err)
	}

	return isExist, nil
}

func (r *rssDomain) createSQLNews(ctx context.Context, tx *sql.Tx, v entity.NewsArticle) (*sql.Tx, entity.NewsArticle, error) {
	// Create a helper function for preparing failure results.
	fail := func(err error) (*sql.Tx, entity.NewsArticle, error) {
		return tx, v, fmt.Errorf("createSQLNews: %v", err)
	}

	_, err := tx.ExecContext(ctx, CreateNews, v.Title, v.URL, v.Content, v.Summary, v.ArticleTS, v.PublishedDate, v.Inserted, v.Updated)
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			// check duplicate constraint
			if mysqlError.Number == 1062 {
				return fail(errors.New("Create New Unique Constraint"))
			}
		}

		return fail(err)
	}

	return tx, v, nil
}
