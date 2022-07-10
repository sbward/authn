package test

import (
	"net/http/httptest"

	app "github.com/keratin/authn"
	"github.com/keratin/authn/server"
)

func Server(app *app.App) *httptest.Server {
	return httptest.NewServer(server.Router(app))
}
