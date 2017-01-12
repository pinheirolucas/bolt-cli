package cache

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var db *bolt.DB

// CreateBoltDB ...
func CreateBoltDB(p string) error {
	var err error

	db, err = bolt.Open(p, 0600, bolt.DefaultOptions)
	if err != nil {
		return errors.Wrap(err, "opening BoltDB")
	}

	return nil
}

// CloseBoltDB ...
func CloseBoltDB() {
	db.Close()
}

// GetDB use only on tests
func GetDB() *bolt.DB {
	return db
}
