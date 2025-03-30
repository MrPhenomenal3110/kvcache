package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"kvcache/internal/cache"
)

// HandlePut handles PUT operations for the cache
func HandlePut(cache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received POST request") // Debugging log
		var req PutRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status:  "ERROR",
				Message: "Invalid request format",
			})
			return
		}
		
		fmt.Println("request body: ")
		fmt.Println(req)
		
		cache.Put(req.Key, req.Value)
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Status:  "OK",
			Message: "Key inserted/updated successfully.",
		})
	}
}

// HandleGet handles GET operations for the cache
func HandleGet(kvCache *cache.Cache) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("Received GET request") // Debugging log

        // Read query parameter
        key := r.URL.Query().Get("key")
        if key == "" {
            http.Error(w, "Missing key parameter", http.StatusBadRequest)
            return
        }

        // Fetch value from cache
        value, found := kvCache.Get(key)

        if !found {
            log.Printf("Key not found: %s", key)
            http.Error(w, "Key not found", http.StatusNotFound)
            return
        }

        fmt.Fprintln(w, value) // Send value to client
    }
}

