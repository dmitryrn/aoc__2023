package main

import (
	"fmt"
	"math"
	"strings"
)

const input = `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func expand(oldMatrix [][]rune) [][]rune {
	matrix := cp(oldMatrix)

	for y := 0; y < len(matrix); y++ {
		hasGals := false
		for x := range matrix[y] {
			if matrix[y][x] > 0 {
				hasGals = true
				break
			}
		}
		if !hasGals {
			before := matrix[0:y]
			after := make([][]rune, len(matrix[y:]))
            copy(after, matrix[y:])

			exp := 1
			between := make([][]rune, exp)
			for i := range between {
				between[i] = make([]rune, len(matrix[0]))
			}

			matrix = append(before, between...)
			matrix = append(matrix, after...)
			y += exp
		}
	}

	return matrix
}

func print(matrix [][]rune) {
	for y := 0; y < len(matrix); y++ {
		fmt.Println(matrix[y])
	}
}

func main() {
	matrix := parseMatrix(input)

	{
		counter := 1
		for y := 0; y < len(matrix); y++ {
			for x := 0; x < len(matrix[0]); x++ {
				if matrix[y][x] == '#' {
					matrix[y][x] = rune(counter)
					counter++
				} else {
					matrix[y][x] = 0
				}
			}
		}
	}

	matrix = expand(matrix)
	matrix = rotate(matrix)
	matrix = expand(matrix)
	matrix = rotate(matrix)
	matrix = rotate(matrix)
	matrix = rotate(matrix)

	gals := [][2]int{}
	{
		counter := 1
		for y := 0; y < len(matrix); y++ {
			for x := 0; x < len(matrix[0]); x++ {
				if matrix[y][x] > 0 {
					gals = append(gals, [2]int{x, y})
					counter++
				}
			}
		}
	}

	distances := 0
	pairs := map[[4]int]struct{}{}
	for _, g1 := range gals {
		for _, g2 := range gals {
			if g1[0] == g2[0] && g1[1] == g2[1] {
				continue
			}
			_, ok := pairs[[4]int{g2[0], g2[1], g1[0], g1[1]}]
			if ok {
				continue
			}
			pairs[[4]int{g1[0], g1[1], g2[0], g2[1]}] = struct{}{}

			d := distance(g1[0], g1[1], g2[0], g2[1])
			distances += d
		}
	}
	fmt.Println(distances)
}

func cp(matrix [][]rune) [][]rune {
	temp := make([][]rune, len(matrix))
	for y := 0; y < len(matrix); y++ {
		temp[y] = make([]rune, len(matrix[y]))
		for x := 0; x < len(matrix[0]); x++ {
			temp[y][x] = matrix[y][x]
		}
	}
	return temp
}

func rotate(matrix [][]rune) [][]rune {
	temp := make([][]rune, len(matrix[0]))
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			temp[x] = append(temp[x], matrix[y][x])
		}
	}
	return temp
}

func distance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x2)-float64(x1)) + math.Abs(float64(y2)-float64(y1)))
}

func parseMatrix(input string) [][]rune {
	lines := strings.Split(input[1:len(input)-1], "\n")

	matrix := [][]rune{}
	for _, l := range lines {
		matrix = append(matrix, []rune(l))
	}

	return matrix
}
