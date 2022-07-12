package server

import (
	"fmt"
	"log"
	"net/http"

	app "github.com/keratin/authn/v2"
)

func Server(app *app.App) {
	if app.Config.PublicPort != 0 {
		go func() {
			fmt.Println(fmt.Sprintf("PUBLIC_PORT: %d", app.Config.PublicPort))
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.PublicPort), PublicRouter(app)))
		}()
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.ServerPort), Router(app)))
}
