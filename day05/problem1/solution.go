package day5_problem1

import (
	"advent2024/day05"
	"advent2024/util"
	"fmt"
)

func calculateMiddleOfCorrectlyOrderedUpdates(rules map[int]util.Set[int], updates [][]int) int {
	sum := 0

	for _, update := range updates {
		if day05.IsCorrectlyOrdered(update, rules) {
			sum += day05.FindMiddle(update)
		}
	}

	return sum
}

type Day5Solution1 struct{}

func (Day5Solution1) Solve(path string) {
	file := util.ReadFile(path)
	rules, updates := day05.ParseInput(file)

	sum := calculateMiddleOfCorrectlyOrderedUpdates(rules, updates)

	fmt.Println(sum)
}
