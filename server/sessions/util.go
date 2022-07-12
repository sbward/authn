package sessions

import (
	"net/http"

	app "github.com/keratin/authn/v2"
	"github.com/keratin/authn/v2/models"
	"github.com/keratin/authn/v2/tokens/sessions"
)

func Get(r *http.Request) *sessions.Claims {
	fn, ok := r.Context().Value(sessionKey(0)).(func() *sessions.Claims)
	if ok {
		return fn()
	}
	return nil
}

func GetAccountID(r *http.Request) int {
	fn, ok := r.Context().Value(accountIDKey(0)).(func() int)
	if ok {
		return fn()
	}
	return 0
}

func Set(cfg *app.Config, w http.ResponseWriter, val string) {
	cookie := &http.Cookie{
		Name:     cfg.SessionCookieName,
		Value:    val,
		Path:     cfg.MountedPath,
		Secure:   cfg.ForceSSL,
		HttpOnly: true,
		SameSite: cfg.SameSiteComputed(),
	}
	if val == "" {
		cookie.MaxAge = -1
	}
	http.SetCookie(w, cookie)
}

func GetRefreshToken(r *http.Request) *models.RefreshToken {
	claims := Get(r)
	if claims != nil {
		token := models.RefreshToken(claims.Subject)
		return &token
	}
	return nil
}
