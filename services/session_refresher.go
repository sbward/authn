package services

import (
	"github.com/pkg/errors"
	app "github.com/sbward/authn"
	"github.com/sbward/authn/data"
	"github.com/sbward/authn/lib/route"
	"github.com/sbward/authn/models"
	"github.com/sbward/authn/ops"
	"github.com/sbward/authn/tokens/identities"
	"github.com/sbward/authn/tokens/sessions"
)

func SessionRefresher(
	refreshTokenStore data.RefreshTokenStore, keyStore data.KeyStore, actives data.Actives, cfg *app.Config, reporter ops.ErrorReporter,
	session *sessions.Claims, accountID int, audience *route.Domain,
) (string, error) {
	// track actives
	if actives != nil {
		err := actives.Track(accountID)
		if err != nil {
			reporter.ReportError(errors.Wrap(err, "Track"))
		}
	}

	// extend refresh token expiration
	err := refreshTokenStore.Touch(models.RefreshToken(session.Subject), accountID)
	if err != nil {
		return "", errors.Wrap(err, "Touch")
	}

	// create new identity token
	identityToken, err := identities.New(cfg, session, accountID, audience.String()).Sign(keyStore.Key())
	if err != nil {
		return "", errors.Wrap(err, "New")
	}

	return identityToken, nil
}
