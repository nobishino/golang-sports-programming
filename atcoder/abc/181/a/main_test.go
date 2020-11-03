package main

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestSolve(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("solve/naive returns same value", prop.ForAll(
		func(v int) bool {
			want := naive(v)
			got := solve(v)
			return got == want
		},
		gen.IntRange(1, 30),
	))

	properties.TestingRun(t)
}
