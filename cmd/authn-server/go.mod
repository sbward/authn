module github.com/sbward/authn/cmd/authn-server

go 1.18

replace github.com/sbward/authn => ../..

replace github.com/sbward/authn/data/sqlite3 => ../../data/sqlite3

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/jmoiron/sqlx v1.3.5
	github.com/pkg/errors v0.9.1
	github.com/sbward/authn v0.0.0-00010101000000-000000000000
	github.com/sbward/authn/data/sqlite3 v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
)

require (
	cloud.google.com/go/compute v1.7.0 // indirect
	github.com/airbrake/gobrake v3.7.4+incompatible // indirect
	github.com/benbjohnson/ego v0.4.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/caio/go-tdigest v3.1.0+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/getsentry/sentry-go v0.13.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/mattn/go-sqlite3 v1.14.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/trustelem/zxcvbn v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d // indirect
	golang.org/x/oauth2 v0.0.0-20220608161450-d0670ef3b1eb // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)
