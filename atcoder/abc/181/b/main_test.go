package main

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestSolve(t *testing.T) {
	properties := gopter.NewProperties(nil)

	// generator
	restrictedValuesGen := gen.IntRange(1, 1e2).Map(func(n int) in {
		elemGen := gen.IntRange(1, 1e6)
		a, b := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			v, _ := elemGen.Sample()
			va := v.(int)
			v, _ = gen.IntRange(va, 1e6).Sample()
			vb := v.(int)
			a[i], b[i] = va, vb
		}
		return in{
			n: n,
			a: a,
			b: b,
		}
	})

	properties.Property("solve returns same value as naive", prop.ForAll(
		func(input in) bool {
			return solve(input) == naive(input)
		},
		restrictedValuesGen,
	))

	properties.TestingRun(t)
}
