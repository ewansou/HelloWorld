package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestCalculator(t *testing.T) {
	g := Goblin(t)
	g.Describe("Calculator", func() {
		g.It("should show addition of two numbers ", func() {
			g.Assert(Add(1, 2))
		})

		g.It("should show subtraction of two numbers", func() {
			g.Assert(Subtract(5, 3)).Equal(2)
		})

		g.It("should show multiplication of two numbers", func() {
			g.Assert(Multiply(5, 6)).Equal(30)
		})

		g.It("should show division of two numbers", func() {
			g.Assert(Divide(10, 2)).Equal(float64(5))
		})
	})
}
