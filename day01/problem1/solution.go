package day1_problem1

import (
	"advent2024/day01"
	"advent2024/util"
	"fmt"
	"sort"
	"strings"
)

func ascending(l []int) (func(i, j int) bool) {
	return (func(i, j int) bool {
		return l[i] < l[j]
	})
}

type Day1Solution1 struct {}

func (Day1Solution1) Solve(path string) {
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")
	left, right := day01.SplitLines(lines)

	sort.Slice(left, ascending(left))
	sort.Slice(right, ascending(right))

	sum := 0
	for i := range left {
		sum += util.Abs(left[i] - right[i])
	}

	fmt.Println(sum)
}