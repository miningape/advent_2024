package day08

import (
	"advent2024/util"
	"unicode"
)

func FindAntennas(lines []string) map[rune][]util.Vector {
	antennas := make(map[rune][]util.Vector)

	for y, line := range lines {
		for x, ch := range line {
			if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
				_, found := antennas[ch]
				if !found {
					antennas[ch] = make([]util.Vector, 0)
				}

				antennas[ch] = append(antennas[ch], util.Vector{ X: x, Y: y })
			}
		}
	}

	return antennas
}

func IsInside(lines []string, position util.Vector) bool {
	if position.X < 0 || position.Y < 0 || position.Y >= len(lines) {
		return false
	}

	return !(position.X >= len(lines[position.Y]))
}