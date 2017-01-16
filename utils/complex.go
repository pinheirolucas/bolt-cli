package utils

import (
	"strings"

	"github.com/pkg/errors"

	"encoding/json"

	"github.com/pinheirolucas/bolt-cli/validator"
)

// FromComplexValue convert a complex string to an map[string]interface{}.
// Where a valid string is something like: [the_key; the_value; another_key; another_value]
func FromComplexValue(cv string) (map[string]interface{}, error) {
	if !validator.IsComplexValue(cv) {
		return nil, errors.New("the provided value is not a complex value")
	}

	mcv := map[string]interface{}{}
	kvArr := toArray(cv)

	for i := 1; i < len(kvArr); i += 2 {
		k := kvArr[i-1]
		v := kvArr[i]

		if validator.IsJSON(v) {
			var parsedJSON interface{}
			err := json.Unmarshal([]byte(v), &parsedJSON)
			if err != nil {
				return nil, errors.Wrap(err, "JSON unmarshal")
			}

			mcv[k] = parsedJSON
		} else {
			mcv[k] = v
		}
	}

	return mcv, nil
}

func toArray(cv string) []string {
	var arr []string

	ncv := strings.TrimPrefix(cv, validator.KVSTART)
	ncv = strings.TrimSuffix(ncv, validator.KVEND)

	arr = strings.Split(ncv, validator.KVDELIM)

	return arr
}
