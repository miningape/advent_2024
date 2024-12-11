package day10

import (
	"advent2024/util"
	"strconv"
	"strings"
)

func ParseTopograph(file string) (map[util.Vector]int, []util.Vector) {
	topograph := make(map[util.Vector]int)
	trailheads := make([]util.Vector, 0)
	lines := strings.Split(file, "\n")

	for y, line := range lines {
		for x, c := range line {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			cell := util.Vector{ X: x, Y: y }
			topograph[cell] = height

			if height == 0 {
				trailheads = append(trailheads, cell)
			}
		}
	}

	return topograph, trailheads
}

func AdjacentDirections() []util.Vector {
	return []util.Vector {
		{ X: 1, Y: 0 },
		{ X: 0, Y: -1 },
		{ X: -1, Y: 0 },
		{ X: 0, Y: 1 },
	}
}
