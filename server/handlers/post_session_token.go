package handlers

import (
	"net/http"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/lib/parse"

	"github.com/sbward/authn/lib/route"
	"github.com/sbward/authn/server/sessions"
	"github.com/sbward/authn/services"
)

func PostSessionToken(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials struct{ Token string }
		if err := parse.Payload(r, &credentials); err != nil {
			WriteErrors(w, err)
			return
		}
		var err error
		var accountID int

		accountID, err = services.PasswordlessTokenVerifier(
			app.AccountStore,
			app.Reporter,
			app.Config,
			credentials.Token,
		)

		if err != nil {
			if fe, ok := err.(services.FieldErrors); ok {
				WriteErrors(w, fe)
				return
			}

			panic(err)
		}

		sessionToken, identityToken, err := services.SessionCreator(
			app.AccountStore, app.RefreshTokenStore, app.KeyStore, app.Actives, app.Config, app.Reporter,
			accountID, route.MatchedDomain(r), sessions.GetRefreshToken(r),
		)
		if err != nil {
			panic(err)
		}

		// Return the signed session in a cookie
		sessions.Set(app.Config, w, sessionToken)

		// Return the signed identity token in the body
		WriteData(w, http.StatusCreated, map[string]string{
			"id_token": identityToken,
		})
	}
}
