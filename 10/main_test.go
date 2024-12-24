package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipe_From(t *testing.T) {
	assert.True(t, seven.Connects(dash, Left))
	assert.True(t, !seven.Connects(dash, Right))

	assert.True(t, j.Connects(f, Up))
	assert.True(t, j.Connects(f, Left))
	assert.True(t, !j.Connects(f, Right))
	assert.True(t, !j.Connects(f, Down))
}

func Test_GetConnectedPipe(t *testing.T) {
	input := `
.....
.S-7.
.|.|.
.L-J.
.....
`

	matrix, _, _ := parseMatrix(input)

	x, y := getConnectedTube(matrix, -1, -1, 1, 1)
	assert.Equal(t, 2, x)
	assert.Equal(t, 1, y)

	x, y = getConnectedTube(matrix, 3, 2, 3, 3)
	assert.Equal(t, 2, x)
	assert.Equal(t, 3, y)
}
