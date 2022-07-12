package handlers

import (
	"net/http"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/server/sessions"
	"github.com/sbward/authn/services"
)

func DeleteSession(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := services.SessionEnder(app.RefreshTokenStore, sessions.GetRefreshToken(r))
		if err != nil {
			app.Reporter.ReportRequestError(err, r)
		}

		sessions.Set(app.Config, w, "")

		w.WriteHeader(http.StatusOK)
	}
}
