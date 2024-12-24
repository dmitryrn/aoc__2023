package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirstNum(t *testing.T) {
	// "
	// abcone2threexyz"
	// xtwone3four"
	// 4nineeightseven2"
	// zoneight234"
	// 7pqrstsixteen"

	n := getFirstNum("o1o2o", strNums)
	assert.Equal(t, 1, n)

	n = getFirstNum("two1nine", strNums)
	assert.Equal(t, 2, n)
	n = getFirstNum(reverse("two1nine"), strNumsReversed)
	assert.Equal(t, 9, n)

	n = getFirstNum("eightwothree", strNums)
	assert.Equal(t, 8, n)

	n = getFirstNum("1eightwothree", strNums)
	assert.Equal(t, 1, n)
	n = getFirstNum(reverse("1eightwothree"), strNumsReversed)
	assert.Equal(t, 3, n)
}
