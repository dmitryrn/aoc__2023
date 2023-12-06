package main

import (
	"math"
	"strconv"
)

var input = []byte(`Time:        61     70     90     66
Distance:   643   1184   1362   1041`)

func main() {
	pt2()
}

func pt2() {
	time, record := parseInputPt2(input)

	(solve(time, record))
}

func solve(t, d uint64) uint64 {
	disc := float64(t*t - 4*d)
	disc = math.Sqrt(float64(disc)) - 2.
	upper := math.Ceil((float64(t) + disc) / 2.)
	lower := math.Floor((float64(t) - disc) / 2.)
	return uint64(upper-lower) + 1
}

func parseInputPt2(input []byte) (uint64, uint64) {
	timeStr := make([]byte, 0, 4*4)
	distanceStr := make([]byte, 0, 4*4)
	first := true
	for i := 0; i < len(input); i++ {
		if input[i] > 47 && input[i] < 58 {
			if first {
				timeStr = append(timeStr, input[i])
			} else {
				distanceStr = append(distanceStr, input[i])
			}
			continue
		}
		if input[i] == 0xa {
			first = false
		}
	}

	time, _ := strconv.ParseUint(string(timeStr), 10, 64)
	distance, _ := strconv.ParseUint(string(distanceStr), 10, 64)

	return time, distance
}
