package day2_problem1

import (
	"advent2024/day02"
	"advent2024/util"
	"fmt"
	"strings"
)

func testDiff(diff []int) bool {
	isPositive := diff[0] >= 0

	for _, item := range diff {
		value := util.Abs(item)
		if value < 1 || value > 3 {
			return false
		}

		if item >= 0 != isPositive {
			return false
		}

		isPositive = item >= 0
	}

	return true
}

type Day2Solution1 struct{}

func (Day2Solution1) Solve(path string) {
	file := util.ReadFile(path)
	reports := strings.Split(file, "\n")

	sum := 0
	for _, report := range reports {
		levels := day02.ParseReport(report)
		diff := day02.DiffLevels(levels)

		if !testDiff(diff) {
			continue
		}

		sum++
	}

	fmt.Println(sum)
}
