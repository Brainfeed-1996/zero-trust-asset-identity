package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/olivierrobert/zero-trust-asset-identity/internal/identity"
	"github.com/olivierrobert/zero-trust-asset-identity/internal/storage"
	"github.com/olivierrobert/zero-trust-asset-identity/internal/metrics"
)

func main() {
	store := storage.NewMemoryStore()
	svc := identity.NewIdentityService(store, "secret-key", 1*time.Hour)

	http.Handle("/metrics", metrics.Handler())

	http.HandleFunc("/issue", func(w http.ResponseWriter, r *http.Request) {
		assetID := r.URL.Query().Get("asset_id")
		if assetID == "" {
			metrics.AuthFailures.Inc()
			http.Error(w, "missing asset_id", http.StatusBadRequest)
			return
		}

		token, err := svc.IssueSVID(assetID)
		if err != nil {
			metrics.AuthFailures.Inc()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		metrics.SVIDIssued.Inc()
		json.NewEncoder(w).Encode(map[string]string{"svid": token})
	})

	fmt.Println("Zero-Trust Asset Identity service starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
