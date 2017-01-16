package utils

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestFromComplexValue(t *testing.T) {
	g := Goblin(t)

	g.Describe("Converts strings from a expected format into a map", func() {
		g.It("Should fail if the provided value is invalid", func() {
			var entryValue = "[a; b; c]"

			_, err := FromComplexValue(entryValue)

			g.Assert(err == nil).IsFalse()
		})

		g.It("Should not decode strings", func() {
			var entryValue = "[name; Lucas; surname; Pinheiro]"

			result, err := FromComplexValue(entryValue)
			if err != nil {
				g.Fail(err)
			}

			name := result["name"]
			surname := result["surname"]

			g.Assert(name).Equal("Lucas")
			g.Assert(surname).Equal("Pinheiro")
		})

		g.It("Should decode numbers", func() {
			var entryValue = "[name; Lucas; age; 20]"

			result, err := FromComplexValue(entryValue)
			if err != nil {
				g.Fail(err)
			}

			age := result["age"]

			g.Assert(age).Equal(float64(20))
		})

		g.It("Should decode JSON", func() {
			var entryValue = "[person; {\"name\": \"Lucas\"}]"

			result, err := FromComplexValue(entryValue)
			if err != nil {
				g.Fail(err)
			}

			person := result["person"].(map[string]interface{})
			name := person["name"]

			g.Assert(name).Equal("Lucas")
		})

		g.It("Should decode array", func() {
			var entryValue = "[ids; [1, 2, 3]]"

			result, err := FromComplexValue(entryValue)
			if err != nil {
				g.Fail(err)
			}

			ids := result["ids"].([]interface{})

			g.Assert(ids[0]).Equal(float64(1))
			g.Assert(ids[1]).Equal(float64(2))
			g.Assert(ids[2]).Equal(float64(3))
		})
	})
}

func TestToArray(t *testing.T) {
	g := Goblin(t)

	g.Describe("Converts strings from a expected format into a map", func() {
		g.It("Should return an array of even values", func() {
			var entryValue = "[a; b; c; d]"

			result := toArray(entryValue)

			resultLength := len(result)
			g.Assert(resultLength == 0).IsFalse()
			g.Assert((len(result) % 2) == 0).IsTrue()
		})

		g.It("Should return an array with the provided values", func() {
			var entryValue = "[a; b; c; d]"

			result := toArray(entryValue)

			g.Assert(result[0]).Equal("a")
			g.Assert(result[1]).Equal("b")
			g.Assert(result[2]).Equal("c")
			g.Assert(result[3]).Equal("d")
		})
	})
}
