package handlers

import (
	"encoding/json"
	"fmt"
	"webscrap/internal/scraper"
	"net/http"
)


func TikTokHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	data, err := scraper.ScrapeTikTokVideo(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scraping TikTok: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}


func FacebookHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	data, err := scraper.ScrapeFacebookPost(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scraping Facebook: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func InstagramHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	data, err := scraper.ScrapeInstagramPost(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scraping Instagram: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
