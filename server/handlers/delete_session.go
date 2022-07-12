package handlers

import (
	"net/http"

	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/server/sessions"
	"github.com/keratin/authn/v2/services"
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
