package day01

import (
	"strconv"
	"strings"
)

func SplitLines(lines []string) ([]int, []int) {
	leftl := []int {}
	rightl := []int {}

	for _, line := range lines {
		items := strings.Split(line, "   ")

		left, err := strconv.Atoi(items[0])
		if err != nil {
			panic(err)
		}

		leftl = append(leftl, left)

		right, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
		}

		rightl = append(rightl, right)
	}

	return leftl, rightl
}
