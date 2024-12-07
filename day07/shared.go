package day07

import (
	"strconv"
	"strings"
)

type Calibration struct {
	Total int
	Parts []int
}

func ParseInput(file string) []Calibration {
	lines := strings.Split(file, "\n")
	parsed := make([]Calibration, 0, len(lines))
	
	for _, line := range lines {
		split := strings.Split(line, ": ")
		total := split[0]
		key, err := strconv.Atoi(total)
		if err != nil {
			panic(err)
		}

		parts := strings.Split(split[1], " ")
		calibration := Calibration { key, make([]int, 0, len(parts)) }
		for _, part := range parts {
			unit, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			calibration.Parts = append(calibration.Parts, unit)
		}

		parsed = append(parsed, calibration)
	}

	return parsed
}