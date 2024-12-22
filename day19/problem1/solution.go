package day19_problem1

import (
	"advent2024/day19"
	"advent2024/util"
	"fmt"
)

func isPossible(memo map[string]bool, locationToTowel map[int][]string, position int, design string) bool {
	if position >= len(design) {
		return true
	}

	rest := design[position:]
	if possible, found := memo[rest]; found {
		return possible
	}

	memo[rest] = false

	towels := locationToTowel[position]
	for _, towel := range towels {
		if isPossible(memo, locationToTowel, position + len(towel), design) {
			memo[rest] = true
		}
	}

	return memo[rest]
}

type Day19Solution1 struct {}

func (Day19Solution1) Solve(path string) {
	file := util.ReadFile(path)

	towels, designs := day19.ParseInput(file)

	count := 0
	for _, design := range designs {
		towelPositions := day19.TowelsInDesign(towels, design)
		positionToTowel := util.InvertListValuedMap(towelPositions)
		memo := make(map[string]bool)

		if isPossible(memo, positionToTowel, 0, design) {
			count++
		}
	}

	fmt.Println(count)
}