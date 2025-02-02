package scraper

import (
	"errors"
	"webscrap/internal/models"
	"github.com/gocolly/colly"
)


func ScrapeFacebookPost(url string) (models.FacebookPost, error) {
	c := colly.NewCollector()

	var post models.FacebookPost

	c.OnHTML("meta[property='og:title']", func(e *colly.HTMLElement) {
		post.Title = e.Attr("content")
	})

	c.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
		post.ImageURL = e.Attr("content")
	})

	c.OnHTML("meta[property='og:description']", func(e *colly.HTMLElement) {
		post.Description = e.Attr("content")
	})

	err := c.Visit(url)
	if err != nil {
		return models.FacebookPost{}, err
	}

	if post.Title == "" {
		return models.FacebookPost{}, errors.New("failed to scrape post")
	}

	return post, nil
}
