package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/services"
)

func GetAccount(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			WriteNotFound(w, "account")
			return
		}

		account, err := services.AccountGetter(app.AccountStore, id)
		if err != nil {
			if _, ok := err.(services.FieldErrors); ok {
				WriteNotFound(w, "account")
				return
			}

			panic(err)
		}

		formattedLastLogin := ""
		if account.LastLoginAt != nil {
			formattedLastLogin = account.LastLoginAt.Format(time.RFC3339)
		}

		formattedPasswordChangedAt := ""
		if !account.PasswordChangedAt.IsZero() {
			formattedPasswordChangedAt = account.PasswordChangedAt.Format(time.RFC3339)
		}

		WriteData(w, http.StatusOK, map[string]interface{}{
			"id":                  account.ID,
			"username":            account.Username,
			"last_login_at":       formattedLastLogin,
			"password_changed_at": formattedPasswordChangedAt,
			"locked":              account.Locked,
			"deleted":             account.DeletedAt != nil,
		})
	}
}
