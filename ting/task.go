package ting

import (
	"github.com/nhk-news-web-easy/nhk-news-fetcher-go"
	"time"
)

const programId = 1

func RunTask() error {
	now := time.Now().UTC()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	newsList, err := nhk_fetcher.FetchNews(startDate, endDate)

	if err != nil {
		return err
	}

	if newsList == nil || len(newsList) == 0 {
		return nil
	}

	dbConfig, err := ParseDbConfig()

	if err != nil {
		return err
	}

	err = InitDb(dbConfig)

	if err != nil {
		return err
	}

	defer CloseDb()

	for _, news := range newsList {
		ting := Ting{
			ProgramId:   programId,
			Title:       news.Title,
			Description: news.OutlineWithRuby,
			AudioUrl:    news.M3u8Url,
			Content:     news.BodyWithoutRuby,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		err = saveTing(ting)

		if err != nil {
			return err
		}
	}

	return nil
}
