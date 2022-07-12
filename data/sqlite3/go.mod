module github.com/sbward/authn/data/sqlite3

go 1.18

replace github.com/sbward/authn => ../../

require (
	github.com/jmoiron/sqlx v1.3.5
	github.com/mattn/go-sqlite3 v1.14.14
	github.com/pkg/errors v0.9.1
	github.com/sbward/authn v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/airbrake/gobrake v3.7.4+incompatible // indirect
	github.com/caio/go-tdigest v3.1.0+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.13.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
