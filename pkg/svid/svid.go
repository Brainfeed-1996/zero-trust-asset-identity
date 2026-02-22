package svid

import (
	"crypto/jwt"
	"time"
)

type SVIDClaims struct {
	SpiffeID string `json:"sub"`
	jwt.StandardClaims
}

func GenerateToken(spiffeID string, key []byte, duration time.Duration) (string, error) {
	claims := SVIDClaims{
		SpiffeID: spiffeID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Issuer:    "zero-trust-asset-identity",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
