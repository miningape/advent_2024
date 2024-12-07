package day7_problem2

import (
	"advent2024/day07"
	"advent2024/util"
	"fmt"
	"strconv"
	"strings"
)

func canListResultInTotal(expected int, parts []int) bool {
	if expected < 0 {
		// We cant subtract or divide a negative number by a positive one and result in a positive number - we can short circuit our search here
		return false
	}

	if len(parts) == 0 {
		return expected == 0
	}

	current := parts[0]
	if next := expected / current;
		expected == next * current && canListResultInTotal(next, parts[1:]) {
		return true
	}

	if next, found := strings.CutSuffix(strconv.Itoa(expected), strconv.Itoa(current)); 
		 found && len(next) != 0 {
		n, err := strconv.Atoi(next)
		if err != nil {
			panic(err)
		}
		
		if canListResultInTotal(n, parts[1:]) {
			return true
		}
	}

	return canListResultInTotal(expected - current, parts[1:])
}

func findPossiblyCorrectCalibrations(calibrations []day07.Calibration) []int {
	correct_calibrations := make([]int, 0)

	for _, calibration := range calibrations {
		if canListResultInTotal(calibration.Total, util.Reverse(calibration.Parts)) {
			correct_calibrations = append(correct_calibrations, calibration.Total)
		}
	}

	return correct_calibrations
}

type Day7Solution2 struct {}

func (Day7Solution2) Solve(path string) {
	file := util.ReadFile(path)
	calibrations := day07.ParseInput(file)

	correct := findPossiblyCorrectCalibrations(calibrations)
	answer := util.SumList(correct)

	fmt.Println(answer)
}
