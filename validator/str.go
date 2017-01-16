package validator

import (
	"strings"
)

const (
	// KVSTART ...
	KVSTART = "["
	// KVEND ...
	KVEND = "]"
	// KVDELIM ...
	KVDELIM = "; "
)

// EndsWith determine if string ends with the suffix
func EndsWith(s, suffix string) bool {
	return strings.LastIndex(s, suffix) == (len(s) - 1)
}

// IsBetween determine if the string starts and ends with the provided prefix and suffix
func IsBetween(s, prefix, suffix string) bool {
	return StartsWith(s, prefix) && EndsWith(s, suffix)
}

// IsComplexValue determine if the string is an array of key value
func IsComplexValue(v string) bool {
	if !IsBetween(v, KVSTART, KVEND) {
		return false
	}

	nv := strings.TrimPrefix(v, KVSTART)
	nv = strings.TrimSuffix(nv, KVEND)

	if (len(strings.Split(nv, KVDELIM)) % 2) > 0 {
		return false
	}

	return true
}

// StartsWith determine if string starts with the prefix
func StartsWith(s, prefix string) bool {
	return strings.Index(s, prefix) == 0
}
