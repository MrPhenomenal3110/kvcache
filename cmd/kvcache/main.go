package main

import (
	"log"
	"net/http"

	"github.com/MrPhenomenal3110/kvcache/internal/api"
	"github.com/MrPhenomenal3110/kvcache/internal/cache"
)

func main() {
	// Initialize cache
	kvCache := cache.NewCache()

	if kvCache == nil {
        log.Fatal("Failed to initialize cache")
    }

	mux := http.NewServeMux()

	// Set up HTTP routes
	mux.HandleFunc("/put", api.HandlePut(kvCache))
	mux.HandleFunc("/get", api.HandleGet(kvCache))
	
	// Start the server
	log.Println("Starting key-value cache server on port 7171...")
	if err := http.ListenAndServe(":7171", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
