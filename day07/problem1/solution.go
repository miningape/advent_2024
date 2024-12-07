package day7_problem1

import (
	"advent2024/util"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(file string) map[int][]int {
	parsed := make(map[int][]int)

	lines := strings.Split(file, "\n")
	for _, line := range lines {
		split := strings.Split(line, ": ")
		total := split[0]
		key, err := strconv.Atoi(total)
		if err != nil {
			panic(err)
		}

		parts := strings.Split(split[1], " ")
		parsed[key] = make([]int, 0, len(parts))
		for _, part := range parts {
			unit, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			parsed[key] = append(parsed[key], unit)
		}
	}

	return parsed
}

func testAllPossibilitesUntil(expected int, parts []int) bool {
	if len(parts) == 1 {
		return expected == parts[0]
	}

	current := parts[0]
	if testAllPossibilitesUntil(expected - current, parts[1:])  {
		return true
	}


	next := expected / current
	if expected != next * current {
		return false
	}
	return testAllPossibilitesUntil(next, parts[1:])
}

func findPossiblyCorrectCalibrations(calibrations map[int][]int) []int {
	correct_calibrations := make([]int, 0)

	for calibration, parts := range calibrations {
		if testAllPossibilitesUntil(calibration, util.Reverse(parts)) {
			correct_calibrations = append(correct_calibrations, calibration)
		}
	}

	return correct_calibrations
}


type Day7Solution1 struct {}

func (Day7Solution1) Solve(path string) {
	file := util.ReadFile(path)
	calibrations := parseInput(file)

	correct := findPossiblyCorrectCalibrations(calibrations)
	answer := util.Sum(correct)

	// fmt.Println(file)
	// fmt.Println()
	// fmt.Println(correct)
	// fmt.Println()
	fmt.Println(answer)
}
