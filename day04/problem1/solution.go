package day4_problem1

import (
	"advent2024/day04"
	"advent2024/util"
	"fmt"
	"strings"
)

func allUnitVectors() []util.Vector {
	vectors := make([]util.Vector, 0, 8)

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}

			vectors = append(vectors, util.Vector{X: x, Y: y})
		}
	}

	return vectors
}

func findXmas(lines []string) int {
	unitVectors := allUnitVectors()
	count := 0

	for y, line := range lines {
		for x := range line {
			for _, direction := range unitVectors {
				if day04.DirectionMatches("XMAS", lines, util.Vector{X: x, Y: y}, direction) {
					count++
				}
			}
		}
	}

	return count
}

type Day4Solution1 struct{}

func (Day4Solution1) Solve(path string) {
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")

	fmt.Println(findXmas(lines))
}
