package validator

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestIsJSON(t *testing.T) {
	g := Goblin(t)

	g.Describe("JSON format validator", func() {
		g.It("Should be valid if the JSON string is correct", func() {
			var jstr = `{
				"name": "Lucas",
				"surname": "Pinheiro"
			}`

			validation := IsJSON(jstr)

			g.Assert(validation).IsTrue()
		})

		g.It("Should be invalid if the JSON string is incorrect", func() {
			var jstr = `{
				"name": "Lucas"
				"surname": "Pinheiro"
			}`

			validation := IsJSON(jstr)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid if the string is not a JSON", func() {
			var jstr = "batata"

			validation := IsJSON(jstr)

			g.Assert(validation).IsFalse()
		})
	})
}
