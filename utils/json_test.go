package utils

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/franela/goblin"
	"github.com/pkg/errors"

	"github.com/pinheirolucas/bolt-cli/mock"
)

func TestRetrieveJSONData(t *testing.T) {
	g := Goblin(t)

	g.Describe("Turn json string into a Go interface{}", func() {
		var tFile *mock.TmpFile

		g.Before(func() {
			var err error

			tFile, err = mock.NewTmpFile()
			if err != nil {
				g.Fail(errors.Wrap(err, "creating temp json file"))
			}
		})

		g.After(func() {
			err := tFile.Remove()
			if err != nil {
				g.Fail(errors.Wrap(err, "removing temp json file"))
			}
		})

		g.It("Should return the right map from JSON file", func() {
			var jstr = `{
	"name": "Lucas",
	"surname": "Pinheiro"
}`

			err := ioutil.WriteFile(tFile.Path, []byte(jstr), os.ModeTemporary)
			if err != nil {
				g.Fail(errors.Wrap(err, "writing temp file fake data"))
			}

			data, err := RetrieveJSONData(tFile.Path)
			if err != nil {
				g.Fail(errors.Wrap(err, "RetrieveJSONData"))
			}

			g.Assert(data["name"]).Equal("Lucas")
			g.Assert(data["surname"]).Equal("Pinheiro")
		})

		g.It("Should return an error if the file does not exists", func() {
			var path = "/a/b/c/batata.json"

			_, err := RetrieveJSONData(path)

			g.Assert(err == nil).IsFalse()
		})

		g.It("Should return an error if the JSON format is not valid", func() {
			var jstr = `{
	"name": "Lucas"
	"surname": "Pinheiro"
}`

			err := ioutil.WriteFile(tFile.Path, []byte(jstr), os.ModeTemporary)
			if err != nil {
				g.Fail(errors.Wrap(err, "writing temp file fake data"))
			}

			_, err = RetrieveJSONData(tFile.Path)

			g.Assert(err == nil).IsFalse()
		})
	})
}
