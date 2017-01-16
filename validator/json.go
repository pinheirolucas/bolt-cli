package validator

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// JSONPathValid determine if the provided JSON path is valid
func JSONPathValid(p string) error {
	fi, err := os.Stat(p)
	if os.IsNotExist(err) {
		return errors.New("the provided path does not exists")
	}

	mode := fi.Mode()
	if mode.IsDir() {
		return errors.New("the provided path is a dir")
	}

	if filepath.Ext(p) != ".json" {
		return errors.New("the provided path is not a JSON file")
	}

	return nil
}

// IsJSON determine if the given string is a valid JSON
func IsJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}
