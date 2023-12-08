package main

import "testing"

func Benchmark(b *testing.B) {
	world, isntructions, start := parseInput(inputPt2Test)

	startCopy := make([]string, len(start))
	copy(startCopy, start)

	for i := 0; i < b.N; i++ {
		pt2(world, isntructions, startCopy)

		copy(startCopy, start)
	}
}
