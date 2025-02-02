package scraper


import (
	"errors"
	"webscrap/internal/models"
	"github.com/gocolly/colly"
)

func ScrapeInstagramPost(url string) (models.InstagramPost, error) {
	c := colly.NewCollector()

	var post models.InstagramPost

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
		return models.InstagramPost{}, err
	}

	
	if post.Title == "" || post.ImageURL == "" {
		return models.InstagramPost{}, errors.New("failed to scrape Instagram post data")
	}

	return post, nil

}