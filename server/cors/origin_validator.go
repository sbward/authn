package cors

import (
	"github.com/sbward/authn/lib/route"
)

func OriginValidator(domains []route.Domain) func(string) bool {
	return func(origin string) bool {
		return route.FindDomain(origin, domains) != nil
	}
}
