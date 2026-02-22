package identity

import (
	"fmt"
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
	if s.Store.IsRevoked(assetID) {
		return "", fmt.Errorf("identity for asset %s is revoked", assetID)
	}
	// In a real scenario, we'd validate assetID against s.Store
	spiffeID := "spiffe://example.org/asset/" + assetID
	return svid.GenerateToken(spiffeID, s.Key, s.Expiry)
}

func (s *IdentityService) RevokeSVID(assetID string) error {
	return s.Store.RevokeIdentity(assetID)
}

func (s *IdentityService) ValidateSVID(assetID string) bool {
	return !s.Store.IsRevoked(assetID)
}
