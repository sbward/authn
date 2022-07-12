package identities

import (
	"strconv"
	"time"

	"github.com/sbward/authn/data/private"

	"github.com/pkg/errors"
	app "github.com/sbward/authn"
	"github.com/sbward/authn/tokens/sessions"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type Claims struct {
	AuthTime *jwt.NumericDate `json:"auth_time"`
	jwt.Claims
}

func (c *Claims) Sign(key *private.Key) (string, error) {
	jwk := jose.JSONWebKey{
		Key:   key.PrivateKey,
		KeyID: key.JWK.KeyID,
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: jwk},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", errors.Wrap(err, "NewSigner")
	}
	return jwt.Signed(signer).Claims(c).CompactSerialize()
}

func New(cfg *app.Config, session *sessions.Claims, accountID int, audience string) *Claims {
	return &Claims{
		AuthTime: session.IssuedAt,
		Claims: jwt.Claims{
			Issuer:   session.Issuer,
			Subject:  strconv.Itoa(accountID),
			Audience: jwt.Audience{audience},
			Expiry:   jwt.NewNumericDate(time.Now().Add(cfg.AccessTokenTTL)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
}
