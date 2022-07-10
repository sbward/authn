package handlers

import (
	"net/http"

	app "github.com/keratin/authn"
	"github.com/keratin/authn/lib/parse"

	"github.com/keratin/authn/lib/route"
	"github.com/keratin/authn/server/sessions"
	"github.com/keratin/authn/services"
)

func PostSession(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials struct {
			Username string
			Password string
		}
		if err := parse.Payload(r, &credentials); err != nil {
			WriteErrors(w, err)
			return
		}

		// Check the password
		account, err := services.CredentialsVerifier(
			app.AccountStore,
			app.Config,
			credentials.Username,
			credentials.Password,
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
			account.ID, route.MatchedDomain(r), sessions.GetRefreshToken(r),
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
