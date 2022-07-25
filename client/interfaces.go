package client

import (
	"github.com/sbward/authn"
	"github.com/sbward/authn/tokens/identities"
	jose "gopkg.in/square/go-jose.v2"
)

// Provides a JSON Web Key from a Key ID
// Wanted to use function signature from go-jose.v2
// but that would make us lose error information
type JWKProvider interface {
	Key(kid string) ([]jose.JSONWebKey, error)
}

// Extracts verified in-built claims from a jwt idToken
type JWTClaimsExtractor interface {
	GetVerifiedClaims(idToken string) (*identities.Claims, error)
}

type appJWKProviderAdapter struct {
	app *authn.App
}

var _ JWKProvider = (*appJWKProviderAdapter)(nil)

func (p *appJWKProviderAdapter) Key(kid string) ([]jose.JSONWebKey, error) {
	set := jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{},
	}
	for _, key := range p.app.KeyStore.Keys() {
		set.Keys = append(set.Keys, key.JWK)
	}
	return set.Key(kid), nil
}
