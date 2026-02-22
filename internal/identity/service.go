package identity

import (
	"time"
	"github.com/olivierrobert/zero-trust-asset-identity/internal/storage"
	"github.com/olivierrobert/zero-trust-asset-identity/pkg/svid"
)

type IdentityService struct {
	Store  storage.IdentityStore
	Key    []byte
	Expiry time.Duration
}

func NewIdentityService(store storage.IdentityStore, key string, expiry time.Duration) *IdentityService {
	return &IdentityService{
		Store:  store,
		Key:    []byte(key),
		Expiry: expiry,
	}
}

func (s *IdentityService) IssueSVID(assetID string) (string, error) {
	// In a real scenario, we'd validate assetID against s.Store
	spiffeID := "spiffe://example.org/asset/" + assetID
	return svid.GenerateToken(spiffeID, s.Key, s.Expiry)
}
