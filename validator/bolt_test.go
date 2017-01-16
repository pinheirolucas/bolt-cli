package validator

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestBoltPathProvided(t *testing.T) {
	g := Goblin(t)

	g.Describe("Check the args", func() {
		g.It("Should return nil if args length is greater than 1", func() {
			var args = []string{"a", "b"}

			validation := boltPathProvided(args)

			g.Assert(validation).Equal(nil)
		})

		g.It("Should return an error if args is empty", func() {
			var args = []string{}

			validation := boltPathProvided(args)

			g.Assert(validation == nil).IsFalse()
		})
	})
}

func TestBoltDBValid(t *testing.T) {
	g := Goblin(t)

	g.Describe("Check the provided DB type", func() {
		g.It("Should return nil if the file extension is '.db'", func() {
			var path = "/abc/bolt.db"

			validation := boltDBValid(path)

			g.Assert(validation).Equal(nil)
		})

		g.It("Should return an error if the file extension does not exists", func() {
			var path = "/abc/bolt"

			validation := boltDBValid(path)

			g.Assert(validation == nil).IsFalse()
		})

		g.It("Should return an error if the file extension is different from '.db'", func() {
			var path = "/abc/bolt.json"

			validation := boltDBValid(path)

			g.Assert(validation == nil).IsFalse()
		})
	})
}
