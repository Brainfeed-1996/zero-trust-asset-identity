package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/olivierrobert/zero-trust-asset-identity/internal/identity"
	"github.com/olivierrobert/zero-trust-asset-identity/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	svc := identity.NewIdentityService(store, "secret-key", 1*time.Hour)

	http.HandleFunc("/issue", func(w http.ResponseWriter, r *http.Request) {
		assetID := r.URL.Query().Get("asset_id")
		if assetID == "" {
			http.Error(w, "missing asset_id", http.StatusBadRequest)
			return
		}

		token, err := svc.IssueSVID(assetID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"svid": token})
	})

	fmt.Println("Zero-Trust Asset Identity service starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
