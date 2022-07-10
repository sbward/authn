package handlers_test

import (
	"net/url"
	"testing"

	"github.com/keratin/authn/lib/route"
	"github.com/keratin/authn/server/test"
	"github.com/keratin/authn/services"
	"github.com/stretchr/testify/require"
)

func TestPostPasswordScore(t *testing.T) {
	test.App()
	app := test.App()
	server := test.Server(app)
	defer server.Close()

	client := route.NewClient(server.URL).Referred(&app.Config.ApplicationDomains[0])

	t.Run("Should successfully score password using JSON request", func(t *testing.T) {
		res, err := client.PostJSON("/password/score", map[string]interface{}{"password": "aSmallPassword"})

		require.NoError(t, err)
		test.AssertData(t, res, map[string]interface{}{"score": 3, "requiredScore": 2})
	})

	t.Run("Should successfully score password using form request", func(t *testing.T) {
		res, err := client.PostForm("/password/score", url.Values{"password": []string{"anotherBetterPassword!"}})

		require.NoError(t, err)
		test.AssertData(t, res, map[string]interface{}{"score": 4, "requiredScore": 2})
	})

	t.Run("Should accuse missing password", func(t *testing.T) {
		res, err := client.PostJSON("/password/score?password=", map[string]interface{}{})

		require.NoError(t, err)
		test.AssertErrors(t, res, services.FieldErrors{services.FieldError{
			Field:   "password",
			Message: services.ErrMissing,
		}})
	})

}
