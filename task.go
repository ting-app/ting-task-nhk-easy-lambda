package main

import (
	"github.com/nhk-news-web-easy/nhk-news-fetcher-go"
	"time"
)

func runTask() error {
	now := time.Now().UTC()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	newsList, err := nhk_fetcher.FetchNews(startDate, endDate)

	if err != nil {
		return err
	}

	for _ = range newsList {

	}

	return nil
}
