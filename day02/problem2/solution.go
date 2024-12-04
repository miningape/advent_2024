package day2_problem2

import (
	"advent2024/day02"
	"advent2024/util"
	"fmt"
	"strings"
)

func remove[T interface{}](index int, list []T) []T {
	next := make([]T, 0, len(list)-1)

	for i, e := range list {
		if i != index {
			next = append(next, e)
		}
	}

	return next
}

func mostCommonSignOnDiff(diff []int) bool {
	count := 0

	for _, e := range diff {
		if e < 0 {
			count--
		}

		if e > 0 {
			count++
		}
	}

	return count >= 0
}

func n(num int) (int, int) {
	if num == 0 {
		return 1, 2
	}

	return num - 1, num + 1
}

func hasCorrectSign(sign bool, num int) bool {
	return (num != 0 && num > 0) == sign
}

func isContrained(num int) bool {
	value := util.Abs(num)
	return value >= 1 && value <= 3
}

func passes(num int, sign bool) bool {
	return num != 0 && hasCorrectSign(sign, num) && isContrained(num)
}

func testDiff(diff []int, levels []int) (bool, int) {
	isPositive := mostCommonSignOnDiff(diff)

	for i, item := range diff {
		if !passes(item, isPositive) {
			left, right := n(i)
			testDiffItem := levels[left] - levels[right]
			if passes(testDiffItem, isPositive) && ((i == len(diff)-1) || passes(diff[i+1], isPositive)) {
				return false, i
			}

			return false, i + 1
		}
	}

	return true, 0
}

func testLevels_(levels []int, hasRemoved bool) bool {
	diff := day02.DiffLevels(levels)
	result, index := testDiff(diff, levels)
	if result {
		return true
	}

	if hasRemoved {
		return false
	}

	newLevels := remove(index, levels)
	return testLevels_(newLevels, true)
}

func testLevels(levels []int) bool {
	return testLevels_(levels, false)
}

type Day2Solution2 struct{}

func (Day2Solution2) Solve(path string) {
	file := util.ReadFile(path)
	reports := strings.Split(file, "\n")

	sum := 0
	for _, report := range reports {
		levels := day02.ParseReport(report)

		if testLevels(levels) {
			sum++
		}
	}

	fmt.Println(sum)
}
