package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dtkachenko/vermilion/internal/storage"
)

func PodsHandler(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pods, err := store.GetAll()
		if err != nil {
			http.Error(w, "Failed to get pods", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pods)

	}

}
