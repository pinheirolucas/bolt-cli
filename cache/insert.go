package cache

import (
	"encoding/json"
	"reflect"

	"strconv"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"

	"github.com/pinheirolucas/bolt-cli/utils"
)

// InsertBucketValue create a bucket with the provided name and add the provided value
func InsertBucketValue(name string, v map[string]interface{}) error {
	return db.Update(func(tx *bolt.Tx) error {
		nameb := []byte(name)

		bucket, err := tx.CreateBucketIfNotExists(nameb)
		if err != nil {
			return errors.Wrap(err, "creating bucket")
		}

		return insertJSONValuesIntoBucket(bucket, v)
	})
}

// InsertComplexValue create a bucket with the provided name.
// The values are represented by a []string
func InsertComplexValue(name string, cv string) error {
	ccv, err := utils.FromComplexValue(cv)
	if err != nil {
		return errors.Wrap(err, "converting from complex value")
	}

	err = InsertBucketValue(name, ccv)
	if err != nil {
		return errors.Wrap(err, "inserting values into bucket")
	}

	return nil
}

// InsertSimpleValue creaate or update some value
func InsertSimpleValue(b, k, v string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bb := []byte(b)
		kb := []byte(k)
		vb := []byte(v)

		bucket, err := tx.CreateBucketIfNotExists(bb)
		if err != nil {
			return err
		}

		return bucket.Put(kb, vb)
	})
}

func insertArray(bkt *bolt.Bucket, k string, v []interface{}) error {
	kb := []byte(k)

	e, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "wrong array format")
	}

	return bkt.Put(kb, e)
}

func insertJSONValuesIntoBucket(bkt *bolt.Bucket, value map[string]interface{}) error {
	for k, v := range value {
		switch reflect.ValueOf(v).Kind() {
		case reflect.Map:
			kb := []byte(k)

			nbkt, err := bkt.CreateBucketIfNotExists(kb)
			if err != nil {
				return errors.Wrap(err, "creating new bucket")
			}

			return insertJSONValuesIntoBucket(nbkt, v.(map[string]interface{}))
		case reflect.String:
			insertString(bkt, k, v.(string))
		case reflect.Float64:
			insertNumber(bkt, k, v.(float64))
		case reflect.Slice:
			insertArray(bkt, k, v.([]interface{}))
		default:
			return errors.New("unsuported type on JSON file")
		}
	}

	return nil
}

func insertNumber(bkt *bolt.Bucket, k string, v float64) error {
	kb := []byte(k)

	vi := int64(v)
	vb := []byte(strconv.FormatInt(vi, 10))

	return bkt.Put(kb, vb)
}

func insertString(bkt *bolt.Bucket, k, v string) error {
	kb := []byte(k)
	vb := []byte(v)

	return bkt.Put(kb, vb)
}
