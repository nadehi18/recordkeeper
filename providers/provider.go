/*
Package providers defines the provider interface which essentially means that
every provider needs to implement GetIP and SetIP functions.  The provider
will then be added as an option in the GetProvider function which returns a
new object of the specified provider.
*/
package providers

import (
	"github.com/nadehi18/recordkeeper/providers/cloudflare"
	"github.com/nadehi18/recordkeeper/record"
)

// Provider interface defines what functions a provider struct should have
type Provider interface {
	GetIP(record.Entry) string
	SetIP(string, record.Entry) bool
}

// GetProvider returns a new object of the given provider type
func GetProvider(name string, username string, authToken string) Provider {

	if name == "cloudflare" {
		return cloudflare.New(username, authToken)
	}

	return &cloudflare.Cloudflare{}
}
