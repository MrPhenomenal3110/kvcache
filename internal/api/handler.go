package api

import (
	"encoding/json"
	"net/http"

	"github.com/MrPhenomenal3110/kvcache/internal/cache"
)

// HandlePut handles PUT operations for the cache
func HandlePut(cache *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
        // Read query parameter
        key := r.URL.Query().Get("key")
        if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status:  "ERROR",
				Message: "Missing key parameter",
			})
            return
        }

        // Fetch value from cache
        value, found := kvCache.Get(key)

        if !found {
            w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Status:  "ERROR",
				Message: "Key not found",
			})
            return
        }

        w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Status:  "OK",
			Key: key,
			Value: value,
		})
    }
}

