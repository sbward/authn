package cors

import (
	"github.com/keratin/authn/lib/route"
)

func OriginValidator(domains []route.Domain) func(string) bool {
	return func(origin string) bool {
		return route.FindDomain(origin, domains) != nil
	}
}
