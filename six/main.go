package main

import (
	"strconv"
	"strings"
)

const input = `
Time:        61     70     90     66
Distance:   643   1184   1362   1041
`

func main() {
	trimmed := strings.TrimSpace(input)

	parts := strings.Split(trimmed, "\n")
	timeLine := parts[0]
	distanceLine := parts[1]

	time := parseLinePt2(timeLine)
	record := parseLinePt2(distanceLine)

	ways := 0

	for i := 1; i < time; i++ {
		speed := i
		remainingTime := time - i
		distanceTraveled := speed * remainingTime

		// println(speed, remainingTime, distanceTraveled)

		if distanceTraveled > record {
			ways++
		}
		if ways > 0 && distanceTraveled < record {
			break
		}
	}

	println(ways)
}

func lineToArr(str string) []int {
	parts := strings.Split(str, " ")
	result := []int{}

	for i := 1; i < len(parts); i++ {
		if len(parts[i]) == 0 {
			continue
		}
		n, err := strconv.Atoi(strings.TrimSpace(parts[i]))
		if err != nil {
			panic(err)
		}

		result = append(result, n)
	}

	return result
}

func parseLinePt2(str string) int {
	parts := strings.Split(str, " ")
	result := ""

	for i := 1; i < len(parts); i++ {
		if len(parts[i]) == 0 {
			continue
		}

		result += parts[i]
	}

	n, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}
	return n
}
