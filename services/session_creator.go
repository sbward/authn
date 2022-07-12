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

func SessionCreator(
	accountStore data.AccountStore, refreshTokenStore data.RefreshTokenStore, keyStore data.KeyStore, actives data.Actives, cfg *app.Config, reporter ops.ErrorReporter,
	accountID int, audience *route.Domain, existingToken *models.RefreshToken,
) (string, string, error) {
	var err error
	err = SessionEnder(refreshTokenStore, existingToken)
	if err != nil {
		reporter.ReportError(errors.Wrap(err, "SessionEnder"))
	}

	// track actives
	if actives != nil {
		err = actives.Track(accountID)
		if err != nil {
			reporter.ReportError(errors.Wrap(err, "Track"))
		}
	}

	// track last activity
	_, err = accountStore.SetLastLogin(accountID)
	if err != nil {
		reporter.ReportError(errors.Wrap(err, "SetLastLogin"))
	}

	// create new session token
	session, err := sessions.New(refreshTokenStore, cfg, accountID, audience.String())
	if err != nil {
		return "", "", errors.Wrap(err, "sessions.New")
	}
	sessionToken, err := session.Sign(cfg.SessionSigningKey)
	if err != nil {
		return "", "", errors.Wrap(err, "session.Sign")
	}

	// create new identity token
	identityToken, err := identities.New(cfg, session, accountID, audience.String()).Sign(keyStore.Key())
	if err != nil {
		return "", "", errors.Wrap(err, "identities.New")
	}

	return sessionToken, identityToken, nil
}
