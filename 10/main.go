package main

import (
	"errors"
	"fmt"
	"strings"
)

const input = `
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`

func main() {
	matrix, startX, startY := parseMatrix(input)

	edges := map[[2]int]struct{}{}

	x := startX
	y := startY
	prevX := -1
	prevY := -1
	count := 0
	for {
		edges[[2]int{x, y}] = struct{}{}
		prevXMemo, prevYMemo := x, y
		x, y = getConnectedTube(matrix, prevX, prevY, x, y)
		if x == startX && y == startY {
			break
		}
		prevX = prevXMemo
		prevY = prevYMemo
		count++
	}
	// println("pt1", (count+1)/2)

	insideCounter := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if _, ok := edges[[2]int{x, y}]; ok { // on edge, skipping
				continue
			}

			counter := 0
			for x := x; x < len(matrix[0]); x++ {
				if _, ok := edges[[2]int{x, y}]; !ok {
					continue
				}

				counter++

				if x == 0 {
					continue
				}
				pipeA, err := Pipe{}.From(matrix[y][x])
				if err != nil {
					panic("expected a pipe")
				}
				pipeB, err := Pipe{}.From(matrix[y][x-1])
				if err == nil {
					if pipeA.Connects(pipeB, Left) {
						counter = 0
					}
				}
			}

			println(x, y, counter)
			if counter%2 != 0 {
				insideCounter++
			}
		}
	}
	println("pt2", insideCounter)
}

func parseMatrix(input string) ([][]byte, int, int) {
	lines := strings.Split(input[1:len(input)-1], "\n")

	startX := -1
	startY := -1
	matrix := [][]byte{}
	for y, l := range lines {
		matrix = append(matrix, []byte(l))

		if startX < 0 && startY < 0 {
			for x, c := range []byte(l) {
				if c == 'S' {
					startY = y
					startX = x
				}
			}
		}
	}

	return matrix, startX, startY
}

func getConnectedTube(matrix [][]byte,
	ignoreX, ignoreY int, startX, startY int,
) (int, int) {
	pipeA, err := Pipe{}.From(matrix[startY][startX])
	if err != nil {
		panic("not a pipe")
	}

	for i, pair := range [4][2]int{
		{startY, startX + 1}, // right
		{startY + 1, startX}, // down
		{startY, startX - 1}, // left
		{startY - 1, startX}, // up
	} {

		y, x := pair[0], pair[1]
		if !inBounds(matrix, x, y) {
			continue
		}
		if x == ignoreX && y == ignoreY {
			continue
		}

		pipeB, err := Pipe{}.From(matrix[y][x])
		if err != nil {
			continue
		}

		// fmt.Println("pipeB", string(pipeB.pipe))
		if pipeA.Connects(pipeB, Direction(i)) {
			// println("connects", string(pipeA.pipe), string(pipeB.pipe), Direction(i), x, y)
			return x, y
		} else {
			// println("doesn't connect", string(pipeA.pipe), string(pipeB.pipe), Direction(i))
		}
	}

	fmt.Println("pipeA", string(pipeA.pipe))
	panic("didn't find connections")
}

type Pipe struct {
	pipe       byte
	directions []Direction
}

func (a Pipe) Connects(b Pipe, bRelDirection Direction) bool {
	for _, da := range a.directions {
		if bRelDirection != da {
			continue
		}
		for _, db := range b.directions {
			if da.IsOpposite(db) {
				return true
			}
		}
	}

	return false
}

var (
	seven = Pipe{'7', []Direction{Left, Down}}
	f     = Pipe{'F', []Direction{Right, Down}}
	j     = Pipe{'J', []Direction{Up, Left}}
	s     = Pipe{'S', []Direction{Left, Down, Up, Right}}
	l     = Pipe{'L', []Direction{Up, Right}}
	pipe  = Pipe{'|', []Direction{Down, Up}}
	dash  = Pipe{'-', []Direction{Right, Left}}
)

func (p Pipe) From(c byte) (Pipe, error) {
	// println("From", string(c))
	if c == '7' {
		return seven, nil
	}
	if c == 'F' {
		return f, nil
	}
	if c == 'J' {
		return j, nil
	}
	if c == 'S' {
		return s, nil
	}
	if c == '|' {
		return pipe, nil
	}
	if c == 'L' {
		return l, nil
	}
	if c == '-' {
		return dash, nil
	}

	return Pipe{}, ErrNotPipe
}

var ErrNotPipe = errors.New("not pipe")

type Direction uint

const (
	Right = 0
	Down  = 1
	Left  = 2
	Up    = 3
)

func (a Direction) IsOpposite(b Direction) bool {
	if a == Left && b == Right {
		return true
	}
	if a == Right && b == Left {
		return true
	}

	if a == Up && b == Down {
		return true
	}
	if a == Down && b == Up {
		return true
	}

	return false
}

func inBounds(matrix [][]byte, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x > len(matrix[0])-1 || y > len(matrix)-1 {
		return false
	}
	return true
}
