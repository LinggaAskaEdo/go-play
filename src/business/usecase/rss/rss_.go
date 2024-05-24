package rss

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"github.com/linggaaskaedo/go-play/src/business/entity"
)

func (r *rssUsecase) GetLatestNews(ctx context.Context) error {
	start := time.Now()

	counter := 0

	doc, err := xmlquery.LoadURL(r.opt.RSSUrl)
	if err != nil {
		r.logger.Error(err)

		return err
	}

	v := entity.NewsArticle{}

	channel := xmlquery.Find(doc, "//item")

	for _, n := range channel {
		if n := n.SelectElement("title"); n != nil {
			title := n.InnerText()
			v.Title = title
		}

		if n := n.SelectElement("link"); n != nil {
			link := n.InnerText()
			v.URL = link

			docDetail, err := htmlquery.LoadURL(link)
			if err != nil {
				r.logger.Error(err)
			}

			docDataDetail := htmlquery.FindOne(docDetail, "//div[@class = 'post-content clearfix']")
			strDocDataDetail := htmlquery.InnerText(docDataDetail)
			strDocDataDetail = strings.TrimSpace(strDocDataDetail)
			strDocDataDetail = strings.ReplaceAll(strDocDataDetail, "\t", "")
			strDocDataDetail = strings.ReplaceAll(strDocDataDetail, "\n", "")

			v.Content = strDocDataDetail
		}

		if n := n.SelectElement("pubDate"); n != nil {
			pubDate := n.InnerText()

			timePub, err := time.Parse(time.RFC1123Z, pubDate)
			if err != nil {
				r.logger.Error(err)

				return err
			}

			v.PublishedDate = sql.NullTime{Time: timePub, Valid: true}

			timestamp := timePub.Unix()
			v.ArticleTS = int64(timestamp)
		}

		v.Inserted = sql.NullTime{Time: time.Now(), Valid: true}

		status, err := r.rss.GetNewsByUrl(ctx, v.URL)
		if err != nil {
			r.logger.Error(err)
		}

		r.logger.Debug("URL: ", v.URL, ", Status: ", status)

		if !status {
			_, err := r.rss.CreateNews(ctx, v)
			if err != nil {
				r.logger.Error(err)

				return err
			}

			counter++
		}
	}

	duration := time.Since(start)

	r.logger.Info(counter, " data added successfully in ", duration.Seconds(), " seconds")

	return nil
}
