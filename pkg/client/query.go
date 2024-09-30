package client

import (
	"net/url"
	"slices"
	"strings"
)

// Query maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Query map
// are case-sensitive.
type Query map[string][]string

// Get retrieves the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Query) Get(key string) string {
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set assigns the key to a single value. It replaces any existing
// values associated with the key.
func (v Query) Set(key, value string) {
	if value != "" {
		v[key] = []string{value}
	}
}

// Add appends a value to the list of values associated with the key.
// It appends to any existing values associated with the key.
func (v Query) Add(key, value string) {
	if value != "" {
		v[key] = append(v[key], value)
	}
}

// Encode converts the Query into a URL-encoded string.
// The resulting string is in "URL encoded" form
// ("bar=baz&foo=quux") with keys sorted alphabetically.
func (v Query) Encode() string {
	if len(v) == 0 {
		return ""
	}

	var buf strings.Builder

	// Get all keys and sort them alphabetically
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	// Encode the keys and values
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)

		for _, v := range vs {
			// Add '&' separator if it's not the first key-value pair
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}

			// Write the key-value pair in URL-encoded format
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}
