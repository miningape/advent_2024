package day4_problem2

import (
	"advent2024/day04"
	"advent2024/util"
	"fmt"
	"strings"
)

func allCornerVectors() []util.Vector {
	return []util.Vector {
		{ X:  1, Y:  1 },
		{ X:  1, Y: -1 },
		{ X: -1, Y: -1},
		{ X: -1, Y:  1},
	}
}

func findXmas(lines []string) int {
	unitVectors := allCornerVectors()
	count := 0
	
	for y, line := range lines {
		for x, ch := range line {
			if ch == 'A' {
				current := util.Vector{ X: x, Y: y }

				for _, direction := range unitVectors {
					start := current.Add(direction)
					if day04.DirectionMatches("MAS", lines, start, direction.Opposite()) {
						next := start.Rotate90(current)
						if (day04.DirectionMatches("MAS", lines, next, direction.RotateOrigin90().Opposite())) {
							count++
						}
					}
				}		
			}
		}
	}

	return count
}

type Day4Solution2 struct {}

func (Day4Solution2) Solve(path string) {
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")

	fmt.Println(findXmas(lines))
}