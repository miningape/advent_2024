package day19_problem2

import (
	"advent2024/day19"
	"advent2024/util"
	"fmt"
)

func numberOfPossibleCombinations(memo map[string]int, locationToTowel map[int][]string, position int, design string) int {
	if position == len(design) {
		return 1
	}

	if position > len(design) {
		return 0
	}

	rest := design[position:]
	if possible, found := memo[rest]; found {
		return possible
	}

	memo[rest] = 0

	towels := locationToTowel[position]
	for _, towel := range towels {
		count := numberOfPossibleCombinations(memo, locationToTowel, position + len(towel), design)
		memo[rest] += count
	}

	return memo[rest]
}

type Day19Solution2 struct {}

func (Day19Solution2) Solve(path string) {
	file := util.ReadFile(path)

	towels, designs := day19.ParseInput(file)

	count := 0
	for _, design := range designs {
		towelPositions := day19.TowelsInDesign(towels, design)
		positionToTowel := util.InvertListValuedMap(towelPositions)
		memo := make(map[string]int)

		c := numberOfPossibleCombinations(memo, positionToTowel, 0, design)
		count += c
	}

	fmt.Println(count)
}