package main

import (
	"strconv"
	"strings"
)

const input = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func main() {
	matrix := inputToMatrix(input)

	sum := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			println(x, y)
			if isNumber(matrix[y][x]) {
				ln := 0
				nstr := make([]rune, 1, 3) // could be on stack if used an array
				nstr[0] = matrix[y][x]
				for i := x + 1; inBounds(len(matrix[0]), len(matrix), x, y) && isNumber(matrix[y][i]); i++ {
					ln = i - x
					nstr = append(nstr, matrix[y][i])
				}
				if isAdjacentToSymbol(matrix, x, y, ln) {
					n, err := strconv.Atoi(string(nstr))
					if err != nil {
						panic("insane")
					}
					sum += n
				}

				x += ln + 1
			}
		}
	}

	println(sum)
}

func inputToMatrix(s string) [][]rune {
	trimmed := strings.TrimSpace(s)

	lines := strings.Split(trimmed, "\n")

	matrix := make([][]rune, len(lines))

	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	return matrix
}

func isAdjacentToSymbol(matrix [][]rune, startX, startY, length int) bool {
	for x := startX - 1; x <= startX+length+1; x++ {
		for y := startY - 1; y < startY+2; y++ {
			if inBounds(len(matrix[0]), len(matrix), x, y) && !isNumber(matrix[y][x]) && matrix[y][x] != 46 {
				return true
			}
		}
	}

	return false
}

func inBounds(width, height, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x > width-1 || y > height-1 {
		return false
	}
	return true
}

func isNumber(c rune) bool {
	return c >= 48 && c < 58
}
