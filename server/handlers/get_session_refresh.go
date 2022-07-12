package handlers

import (
	"net/http"

	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/lib/route"
	"github.com/keratin/authn/v2/server/sessions"
	"github.com/keratin/authn/v2/services"
	"github.com/pkg/errors"
)

func GetSessionRefresh(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check for valid session with live token
		accountID := sessions.GetAccountID(r)
		if accountID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		identityToken, err := services.SessionRefresher(
			app.RefreshTokenStore, app.KeyStore, app.Actives, app.Config, app.Reporter,
			sessions.Get(r), accountID, route.MatchedDomain(r),
		)
		if err != nil {
			panic(errors.Wrap(err, "IdentityForSession"))
		}

		WriteData(w, http.StatusCreated, map[string]string{
			"id_token": identityToken,
		})
	}
}
