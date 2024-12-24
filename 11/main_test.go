package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_expand(t *testing.T) {
	t.Run("X", func(t *testing.T) {
		m := parseMatrix(`
#...
...#
`)
		gals := getGals(m)
		gals = expand(gals, 0, 2)
		assert.ElementsMatch(t, [][2]int{{0, 0}, {5, 1}}, gals)

		m = parseMatrix(`
#...
...#
`)
		gals = getGals(m)
		gals = expand(gals, 0, 4)
		assert.ElementsMatch(t, [][2]int{{0, 0}, {9, 1}}, gals)

		m = parseMatrix(`
#..
..#
`)
		gals = getGals(m)
		gals = expand(gals, 0, 2)
		assert.ElementsMatch(t, [][2]int{{0, 0}, {3, 1}}, gals)

		m = parseMatrix(`
#......
...#..#
`)
		gals = getGals(m)
		gals = expand(gals, 0, 2)
		assert.ElementsMatch(t, [][2]int{{0, 0}, {5, 1}, {10, 1}}, gals)
	})

	// Y
	t.Run("Y", func(t *testing.T) {
		// 		m := parseMatrix(`
		// #..
		// ..#
		// `)
		//         gals := getGals(m)
		//         gals = expand(gals, 1, 2)
		// 		assert.ElementsMatch(t, [][2]int{{0, 0}, {2, 1}}, gals)

		m := parseMatrix(`
#.
..
.#
`)
		gals := getGals(m)
		gals = expand(gals, 1, 2)
		assert.ElementsMatch(t, [][2]int{{0, 0}, {1, 3}}, gals)

		m = parseMatrix(`
#
.
.
#
.
.
#
`)
		a := `
#
.
.
.
.
#
.
.
.
.
#
`
		println(a)
		gals = getGals(m)
		gals = expand(gals, 1, 2)
		assert.ElementsMatch(t, [][2]int{{0, 0}, {0, 5}, {0, 10}}, gals)
	})
}
