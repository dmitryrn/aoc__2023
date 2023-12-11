package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_expand(t *testing.T) {
	m := parseMatrix(`
#...
...#
`)
	gals := getGals(m)
	gals = expand(gals, 0, 2)
	assert.ElementsMatch(t, [][2]int{{0, 0}, {5, 1}}, gals)
}
