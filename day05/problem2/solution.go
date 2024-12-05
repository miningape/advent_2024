package day5_problem2

import (
	"advent2024/day05"
	"advent2024/util"
	"fmt"
	"sort"
)

func calculateMiddleOfIncorrectlyOrderedUpdates(rules map[int]util.Set[int], updates [][]int) int {
	sum := 0

	for _, update := range updates {
		if day05.IsCorrectlyOrdered(update, rules) {
			continue
		}
		
		sort.Slice(update, func(i, j int) bool {
			left := update[i]
			right := update[j]
			rule, found := rules[left]

			if found && rule.Contains(right) {
				return true
			}

			return false
		})

		sum += day05.FindMiddle(update)
	}

	return sum
}

type Day5Solution2 struct {}

func (Day5Solution2) Solve(path string) {
	file := util.ReadFile(path)
	rules, updates := day05.ParseInput(file)

	sum := calculateMiddleOfIncorrectlyOrderedUpdates(rules, updates)

	fmt.Println(sum)
}