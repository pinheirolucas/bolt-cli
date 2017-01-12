package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// RetrieveJSONData convert JSON to a generic struct (interface{})
func RetrieveJSONData(p string) (map[string]interface{}, error) {
	file, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, errors.Wrap(err, "opening JSON file")
	}

	var parsedJSON interface{}
	err = json.NewDecoder(bytes.NewBuffer(file)).Decode(&parsedJSON)
	if err != nil {
		return nil, errors.Wrap(err, "decoding JSON file")
	}

	return parsedJSON.(map[string]interface{}), nil
}
