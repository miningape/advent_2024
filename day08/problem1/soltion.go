package day8_problem1

import (
	"advent2024/day08"
	"advent2024/util"
	"fmt"
	"strings"
)

func findAntiNodes(lines []string, antennas map[rune][]util.Vector) util.Set[util.Vector] {
	antiNodes := util.SetOf[util.Vector]()

	for _, antennas := range antennas {
		for i, source := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				other := antennas[j]
				if i == j {
					continue
				}

				direction := other.Sub(source)
				first := source.Add(direction.Mul(2))
				if day08.IsInside(lines, first) {
					antiNodes.Add(first)
				}

				second := source.Sub(direction)
				if day08.IsInside(lines, first) {
					antiNodes.Add(second)
				}
			}
		}
	}

	return antiNodes
}

type Day8Solution1 struct {}

func (Day8Solution1) Solve(path string) {
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")

	antennas := day08.FindAntennas(lines)
	antiNodes := findAntiNodes(lines, antennas)

	fmt.Println(antiNodes)
	fmt.Println("Answer:", len(antiNodes))
}