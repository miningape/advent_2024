package day04

import "advent2024/util"

func isWithin(lines []string, position util.Vector) bool {
	if position.Y < 0 || position.Y >= len(lines) {
		return false
	}

	line := lines[position.Y]
	if position.X < 0 || position.X >= len(line) {
		return false
	}

	return true
}

func get(lines []string, position util.Vector) rune {
	return rune(lines[position.Y][position.X])
}

func DirectionMatches(s string, lines []string, position util.Vector, direction util.Vector) bool {
	next := position

	for _, c := range s {
		if !isWithin(lines, next) {
			return false
		}

		if c != get(lines, next) {
			return false
		}

		next = next.Add(direction)
	}

	return true
}
