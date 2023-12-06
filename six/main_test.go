package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_parseInputPt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseInputPt2([]byte(input))
	}
}

func Benchmark_pt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pt2()
	}
}

// tests

func Test_solve(t *testing.T) {
	assert.Equal(t, uint64(71503), solve(71530, 940200))
}

func Test_parseInputPt2(t *testing.T) {
	time, d := parseInputPt2([]byte(`Time:        61     70     90     66
Distance:   643   1184   1362   1041`))

	assert.Equal(t, uint64(61709066), time)
	assert.Equal(t, uint64(643118413621041), d)
}
