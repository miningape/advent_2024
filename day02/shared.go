package day02

import (
	"strconv"
	"strings"
)

func ParseReport(report string) []int {
	levels := strings.Split(report, " ")
	parsed := make([]int, 0, len(levels))

	for _, level := range levels {
		number, err := strconv.Atoi(level)
		if err != nil {
			panic(err)
		}

		parsed = append(parsed, number)
	}

	return parsed
}

func DiffLevels(levels []int) []int {
	diff := make([]int, 0, len(levels)-1)

	previous := 0
	for i, level := range levels {
		if i != 0 {
			diff = append(diff, previous-level)
		}

		previous = level
	}

	return diff
}
