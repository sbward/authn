package services

import (
	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/data"
	"github.com/keratin/authn/v2/lib/route"
	"github.com/keratin/authn/v2/models"
	"github.com/keratin/authn/v2/ops"
	"github.com/keratin/authn/v2/tokens/identities"
	"github.com/keratin/authn/v2/tokens/sessions"
	"github.com/pkg/errors"
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
