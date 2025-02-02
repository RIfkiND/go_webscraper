package main

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
	"webscrap/internal/handlers"

	"golang.org/x/time/rate"
)

func limitMiddleware(next http.Handler, rLimiter *rate.Limiter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rLimiter.Allow() {
			http.Error(w, "Too many requests. Please try again later.", http.StatusTooManyRequests)
			return
		}

		
		randomDelay := time.Duration(rand.Intn(5)+1) * time.Second 
		time.Sleep(randomDelay)

		next.ServeHTTP(w, r)
	})
}

func main() {
	rand.Seed(time.Now().UnixNano()) 

	
	rLimiter := rate.NewLimiter(1, 5)


	var wg sync.WaitGroup


	endpoints := map[string]http.Handler{
		"/api/v1/scrape/tiktok":    limitMiddleware(http.HandlerFunc(handlers.TikTokHandler), rLimiter),
		"/api/v1/scrape/facebook": limitMiddleware(http.HandlerFunc(handlers.FacebookHandler), rLimiter),
		"/api/v1/scrape/instagram": limitMiddleware(http.HandlerFunc(handlers.InstagramHandler), rLimiter),
	}


	wg.Add(1)
	go func() {
		defer wg.Done()
		for route, handler := range endpoints {
			http.Handle(route, handler)
		}
		port := ":8080"
		log.Printf("Server running on port %s", port)
		log.Fatal(http.ListenAndServe(port, nil))
	}()

	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Println("Background task running...")
			time.Sleep(10 * time.Second) 
		}
	}()

	
	wg.Wait()
}
