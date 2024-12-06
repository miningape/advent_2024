package day5_problem2

import (
	"advent2024/day05"
	"advent2024/util"
	"fmt"
	"sort"
)

func calculateMiddleOfIncorrectlyOrderedUpdate(rules map[int]util.Set[int], update []int) int {
	if day05.IsCorrectlyOrdered(update, rules) {
		return 0
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

	return day05.FindMiddle(update)
}

type Day5Solution2 struct{}

func (Day5Solution2) Solve(path string) {
	file := util.ReadFile(path)
	rules, updates := day05.ParseInput(file)

	sum := 0
	for _, update := range updates {
		sum += calculateMiddleOfIncorrectlyOrderedUpdate(rules, update)
	}

	fmt.Println(sum)
}
