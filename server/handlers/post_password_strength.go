package handlers

import (
	"net/http"

	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/lib/parse"

	"github.com/keratin/authn/v2/services"
)

func PostPasswordScore(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials struct {
			Password string
		}
		if err := parse.Payload(r, &credentials); err != nil {
			WriteErrors(w, err)
			return
		}

		if credentials.Password == "" {
			WriteErrors(w, services.FieldErrors{services.FieldError{
				Field:   "password",
				Message: services.ErrMissing,
			}})
			return
		}

		score := services.CalculatePasswordScore(credentials.Password)

		WriteData(w, http.StatusOK, map[string]interface{}{
			"score":         score,
			"requiredScore": app.Config.PasswordMinComplexity,
		})
	}
}
