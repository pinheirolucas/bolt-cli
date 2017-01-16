package validator

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestEndsWith(t *testing.T) {
	g := Goblin(t)

	g.Describe("String end checker", func() {
		var fullString = "[a b c]"

		g.It("Should be valid when the suffix matches", func() {
			var suffix = "]"

			validation := EndsWith(fullString, suffix)

			g.Assert(validation).IsTrue()
		})

		g.It("Should be invalid when the suffix does not match", func() {
			var suffix = "batata"

			validation := EndsWith(fullString, suffix)

			g.Assert(validation).IsFalse()
		})
	})
}

func TestIsBetween(t *testing.T) {
	g := Goblin(t)

	g.Describe("String prefix and suffix validator", func() {
		var fullString = "[a b c]"

		g.It("Should be valid when prefix and suffix matches", func() {
			var prefix = "["
			var suffix = "]"

			validation := IsBetween(fullString, prefix, suffix)

			g.Assert(validation).IsTrue()
		})

		g.It("Should be invalid when prefix does not match", func() {
			var prefix = "batata"
			var suffix = "]"

			validation := IsBetween(fullString, prefix, suffix)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid when suffix does not match", func() {
			var prefix = "["
			var suffix = "batata"

			validation := IsBetween(fullString, prefix, suffix)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid when prefix and suffix does not match", func() {
			var prefix = "batata"
			var suffix = "batata"

			validation := IsBetween(fullString, prefix, suffix)

			g.Assert(validation).IsFalse()
		})
	})
}

func TestIsComplexValue(t *testing.T) {
	g := Goblin(t)

	g.Describe("String array representation validator", func() {
		g.It("Should be valid with simple values", func() {
			var value = "[name; Lucas; surname; Pinheiro]"

			validation := IsComplexValue(value)

			g.Assert(validation).IsTrue()
		})

		g.It("Should be valid with JSON values", func() {
			var value = "[person; {\"name\": \"Lucas\"}]"

			validation := IsComplexValue(value)

			g.Assert(validation).IsTrue()
		})

		g.It("Should be valid with JSON values", func() {
			var value = "[person; {\"name\": \"Lucas\"}]"

			validation := IsComplexValue(value)

			g.Assert(validation).IsTrue()
		})

		g.It("Should be invalid if does not starts with [", func() {
			var value = "name; Lucas; surname; Pinheiro]"

			validation := IsComplexValue(value)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid if does not ends with ]", func() {
			var value = "[name; Lucas; surname; Pinheiro"

			validation := IsComplexValue(value)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid if does not starts with [ and ends with ]", func() {
			var value = "name; Lucas; surname; Pinheiro"

			validation := IsComplexValue(value)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid if the delims are not '; '", func() {
			var value = "[name Lucas surname Pinheiro]"

			validation := IsComplexValue(value)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid if the array representation is empty", func() {
			var value = "[]"

			validation := IsComplexValue(value)

			g.Assert(validation).IsFalse()
		})

		g.It("Should be invalid if the array representation length is odd", func() {
			var value = "[a; b; c]"

			validation := IsComplexValue(value)

			g.Assert(validation).IsFalse()
		})
	})
}

func TestStartsWith(t *testing.T) {
	g := Goblin(t)

	g.Describe("String start checker", func() {
		var fullString = "[a b c]"

		g.It("Should be valid when the prefix matches", func() {
			var prefix = "["

			validation := StartsWith(fullString, prefix)

			g.Assert(validation).IsTrue()
		})

		g.It("Should be invalid when the prefix does not match", func() {
			var prefix = "batata"

			validation := StartsWith(fullString, prefix)

			g.Assert(validation).IsFalse()
		})
	})
}
