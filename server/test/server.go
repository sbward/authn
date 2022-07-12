package test

import (
	"net/http/httptest"

	app "github.com/sbward/authn"
	"github.com/sbward/authn/server"
)

func Server(app *app.App) *httptest.Server {
	return httptest.NewServer(server.Router(app))
}
