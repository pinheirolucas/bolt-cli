package cache

import (
	"encoding/json"
	"testing"

	"github.com/boltdb/bolt"
	. "github.com/franela/goblin"
	"github.com/pkg/errors"

	"github.com/pinheirolucas/bolt-cli/mock"
)

func TestInsertSimpleValue(t *testing.T) {
	g := Goblin(t)

	g.Describe("Insert a simple value into a bucket", func() {
		var db *bolt.DB
		var tFile *mock.TmpFile

		g.Before(func() {
			var err error

			tFile, err = mock.NewTmpFile()
			if err != nil {
				g.Fail(errors.Wrap(err, "new db path"))
			}

			CreateBoltDB(tFile.Path)
			db = GetDB()
		})

		g.After(func() {
			defer tFile.Remove()
			CloseBoltDB()
		})

		g.It("Should save the value into the provided key", func() {
			var bucketName = "TestBucket"
			var keyName = "name"
			var value = "Lucas"

			var savedValue string

			err := InsertSimpleValue(bucketName, keyName, value)
			if err != nil {
				g.Fail(errors.Wrap(err, "insert simple value"))
			}

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(bucketName))

				savedValue = string(b.Get([]byte(keyName)))

				return nil
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "get inserted value"))
			}

			g.Assert(value).Equal(savedValue)
		})
	})
}

func TestInsertArray(t *testing.T) {
	g := Goblin(t)

	g.Describe("Insert array into bucket", func() {
		var db *bolt.DB
		var tFile *mock.TmpFile

		g.Before(func() {
			var err error

			tFile, err = mock.NewTmpFile()
			if err != nil {
				g.Fail(errors.Wrap(err, "new db path"))
			}

			CreateBoltDB(tFile.Path)
			db = GetDB()
		})

		g.After(func() {
			defer tFile.Remove()
			CloseBoltDB()
		})

		g.It("Should insert an array as encoded JSON with simple values", func() {
			var bucketName = "TestBucket"
			var keyName = "names"
			var names = []interface{}{"Lucas", "Nathan"}

			var savedNames []interface{}

			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
				if err != nil {
					return errors.Wrap(err, "create bucket")
				}

				return insertArray(b, keyName, names)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "set inserted value"))
			}

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(bucketName))

				n := b.Get([]byte(keyName))

				return json.Unmarshal(n, &savedNames)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "get inserted value"))
			}

			g.Assert(savedNames[0]).Equal("Lucas")
			g.Assert(savedNames[1]).Equal("Nathan")
		})

		g.It("Should insert an array as encoded JSON with complex values", func() {
			var bucketName = "TestBucket"
			var keyName = "names"
			var people = []interface{}{
				map[string]interface{}{
					"name": "Lucas",
					"age":  20,
				},
				"Nathan",
			}

			var savedPeople []interface{}

			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
				if err != nil {
					return errors.Wrap(err, "create bucket")
				}

				return insertArray(b, keyName, people)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "set inserted value"))
			}

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(bucketName))

				n := b.Get([]byte(keyName))

				return json.Unmarshal(n, &savedPeople)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "get inserted value"))
			}

			person1 := savedPeople[0].(map[string]interface{})
			person2 := savedPeople[1].(string)

			g.Assert(person1["name"]).Equal("Lucas")
			g.Assert(person1["age"]).Equal(float64(20))

			g.Assert(person2).Equal("Nathan")
		})
	})
}

func TestInsertJSONValuesIntoBucket(t *testing.T) {
	g := Goblin(t)

	g.Describe("Insert map into bucket recursively", func() {
		var db *bolt.DB
		var tFile *mock.TmpFile

		g.Before(func() {
			var err error

			tFile, err = mock.NewTmpFile()
			if err != nil {
				g.Fail(errors.Wrap(err, "new db path"))
			}

			CreateBoltDB(tFile.Path)
			db = GetDB()
		})

		g.After(func() {
			defer tFile.Remove()
			CloseBoltDB()
		})

		g.It("Should insert all the simple values inside of the base bucket", func() {
			var bucketName = "BaseBucket"
			var values = map[string]interface{}{
				"name": "Lucas",
				"age":  float64(20),
			}

			var savedName string
			var savedAge string

			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
				if err != nil {
					return errors.Wrap(err, "create bucket")
				}

				return insertJSONValuesIntoBucket(b, values)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "set inserted value"))
			}

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(bucketName))

				savedName = string(b.Get([]byte("name")))
				savedAge = string(b.Get([]byte("age")))

				return nil
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "get inserted value"))
			}

			g.Assert(savedName).Equal("Lucas")
			g.Assert(savedAge).Equal("20")
		})

		g.It("Should insert all the simple values of the map inside of the base bucket and the nested maps into nested buckets", func() {
			var bucketName = "BaseBucket"
			var values = map[string]interface{}{
				"name": "Lucas",
				"age":  float64(20),
				"father": map[string]interface{}{
					"name": "Sandro",
					"age":  float64(45),
				},
			}

			var savedName string
			var savedAge string
			var fatherName string
			var fatherAge string

			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
				if err != nil {
					return errors.Wrap(err, "create bucket")
				}

				return insertJSONValuesIntoBucket(b, values)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "set inserted value"))
			}

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(bucketName))
				sb := b.Bucket([]byte("father"))

				savedName = string(b.Get([]byte("name")))
				savedAge = string(b.Get([]byte("age")))

				fatherName = string(sb.Get([]byte("name")))
				fatherAge = string(sb.Get([]byte("age")))

				return nil
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "get inserted value"))
			}

			g.Assert(savedName).Equal("Lucas")
			g.Assert(savedAge).Equal("20")

			g.Assert(fatherName).Equal("Sandro")
			g.Assert(fatherAge).Equal("45")
		})
	})
}

func TestInsertNumber(t *testing.T) {
	g := Goblin(t)

	g.Describe("Insert number into bucket", func() {
		var db *bolt.DB
		var tFile *mock.TmpFile

		g.Before(func() {
			var err error

			tFile, err = mock.NewTmpFile()
			if err != nil {
				g.Fail(errors.Wrap(err, "new db path"))
			}

			CreateBoltDB(tFile.Path)
			db = GetDB()
		})

		g.After(func() {
			defer tFile.Remove()
			CloseBoltDB()
		})

		g.It("Should insert a number as string", func() {
			var bucketName = "TestBucket"
			var keyName = "age"
			var n = float64(20)

			var savedN string

			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
				if err != nil {
					return errors.Wrap(err, "create bucket")
				}

				return insertNumber(b, keyName, n)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "set inserted value"))
			}

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(bucketName))

				savedN = string(b.Get([]byte(keyName)))

				return nil
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "set inserted value"))
			}

			g.Assert(savedN).Equal("20")
		})
	})
}

func TestInsertString(t *testing.T) {
	g := Goblin(t)

	g.Describe("Insert string into bucket", func() {
		var db *bolt.DB
		var tFile *mock.TmpFile

		g.Before(func() {
			var err error

			tFile, err = mock.NewTmpFile()
			if err != nil {
				g.Fail(errors.Wrap(err, "new db path"))
			}

			CreateBoltDB(tFile.Path)
			db = GetDB()
		})

		g.After(func() {
			defer tFile.Remove()
			CloseBoltDB()
		})

		g.It("Should insert a string", func() {
			var bucketName = "TestBucket"
			var keyName = "name"
			var name = "Lucas"

			var savedName string

			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
				if err != nil {
					return errors.Wrap(err, "create bucket")
				}

				return insertString(b, keyName, name)
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "set inserted value"))
			}

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(bucketName))

				savedName = string(b.Get([]byte(keyName)))

				return nil
			})
			if err != nil {
				g.Fail(errors.Wrap(err, "get inserted value"))
			}

			g.Assert(savedName).Equal("Lucas")
		})
	})
}
