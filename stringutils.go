package main

import (
	"strings"
)

// stringToMap splits `s string` using `sep` as separator and
// set every word as a key in a new map. The value of
// all keys is set to `true`
func stringToMap(s string, sep string) map[string]bool {
	ss := strings.Split(s, sep)
	m := make(map[string]bool)
	for _, word := range ss {
		m[word] = true
	}

	return m

}
