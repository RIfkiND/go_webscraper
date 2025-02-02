package scraper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"webscrap/internal/models"
)


func ScrapeTikTokVideo(url string) (models.TikTokVideo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.TikTokVideo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.TikTokVideo{}, errors.New("failed to fetch TikTok video")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.TikTokVideo{}, err
	}

	// Extract JSON metadata
	re := regexp.MustCompile(`window\.__INIT_PROPS__ = (.*?);</script>`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return models.TikTokVideo{}, errors.New("failed to extract metadata")
	}

	// Parse JSON
	var rawData map[string]interface{}
	if err := json.Unmarshal([]byte(matches[1]), &rawData); err != nil {
		return models.TikTokVideo{}, err
	}

	// Extract data
	videoData := rawData["VideoPage"].(map[string]interface{})["props"].(map[string]interface{})["pageProps"].(map[string]interface{})["videoData"].(map[string]interface{})

	return models.TikTokVideo{
		Username:    videoData["author"].(map[string]interface{})["uniqueId"].(string),
		Caption:     videoData["desc"].(string),
		VideoURL:    videoData["video"].(map[string]interface{})["playAddr"].(string),
		LikeCount:   int(videoData["stats"].(map[string]interface{})["diggCount"].(float64)),
		CommentCount: int(videoData["stats"].(map[string]interface{})["commentCount"].(float64)),
	}, nil
}
