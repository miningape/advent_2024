package day1_problem2

import (
	"advent2024/day01"
	"advent2024/util"
	"fmt"
	"strings"
)

func countAppearances(slice []int) map[int]int {
	m := make(map[int]int)

	for _, elem := range slice {
		count := m[elem]
		m[elem] = count + 1
	}

	return m
}

type Day1Solution2 struct {}

func (Day1Solution2) Solve(path string) {
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")
	left, right := day01.SplitLines(lines)

	tally := countAppearances(right)

	sum := 0
	for _, item := range left {
		appearances := tally[item]
		sum += item * appearances
	}

	fmt.Println(sum)
}