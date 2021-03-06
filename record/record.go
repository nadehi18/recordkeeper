/*
Package record contains the generic record entry type to store commonly used
information from a given DNS record.

It exports a New function which creates a new entry object.
*/
package record

// Entry is a generic struct that keeps information about a certain record
type Entry struct {
	Domain         string `json:"name" mapstructure:"name"`
	Address        string `json:"content"`
	AddressSetting string `mapstructure:"address"`
	ID             string `json:"-" mapstructure:"ID"`
	ZoneID         string `json:"-" mapstructure:"zoneID"`
	RecordType     string `json:"type" mapstructure:"type"`
	Proxied        bool   `json:"proxied" mapstructure:"proxied"`
	TTL            int    `json:"ttl" mapstructure:"ttl"`
}

// New returns a new record entry
func New() *Entry {
	return &Entry{}
}
