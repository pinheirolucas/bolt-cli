package cache

import (
	"log"
	"reflect"

	"strconv"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

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

// InsertBucketValue create a bucket with the provided name and add the provided value
func InsertBucketValue(name string, v interface{}) error {
	iv := v.(map[string]interface{})

	return db.Update(func(tx *bolt.Tx) error {
		nameb := []byte(name)

		bucket, err := tx.CreateBucketIfNotExists(nameb)
		if err != nil {
			return errors.Wrap(err, "creating bucket")
		}

		return insertJSONValuesIntoBucket(bucket, iv)
	})
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
			return insertString(bkt, k, v.(string))
		case reflect.Float64:
			return insertNumber(bkt, k, v.(float64))
		case reflect.Slice:
			log.Println("Array:", v)
		default:
			return errors.New("unsuported type on JSON file")
		}
	}

	return nil
}

func insertString(bkt *bolt.Bucket, k, v string) error {
	kb := []byte(k)
	vb := []byte(v)

	return bkt.Put(kb, vb)
}

func insertNumber(bkt *bolt.Bucket, k string, v float64) error {
	kb := []byte(k)

	vi := int64(v)
	vb := []byte(strconv.FormatInt(vi, 10))

	return bkt.Put(kb, vb)
}
