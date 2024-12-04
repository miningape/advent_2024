package day4_problem1

import (
	"advent2024/util"
	"fmt"
	"strings"
)

type vector struct {
	x, y int
}

func (left vector) add(right vector) vector {
	return vector{
		x: left.x + right.x,
		y: left.y + right.y,
	}
}

func allUnitVectors() []vector {
	vectors := make([]vector, 0, 8)

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}

			vectors = append(vectors, vector{ x, y })
		}
	}

	return vectors
}

func isWithin(lines []string, position vector) bool {
	if position.y < 0 || position.y >= len(lines) {
		return false
	}

	line := lines[position.y]
	if position.x < 0 || position.x >= len(line) {
		return false
	}

	return true
}

func get(lines []string, position vector) rune {
	return rune(lines[position.y][position.x])
}

func directionMatches(s string, lines []string, position vector, direction vector) bool {
	next := position

	for _, c := range s {
		if !isWithin(lines, next) {
			return false
		}

		if c != get(lines, next) {
			return false
		}

		next = next.add(direction)
	}

	return true
}


func findXmas(lines []string) int {
	unitVectors := allUnitVectors()
	count := 0
	
	for y, line := range lines {
		for x := range line {
			for _, direction := range unitVectors {
				if directionMatches("XMAS", lines, vector{ x, y }, direction) {
					count++
				}
			}
		}
	}

	return count
}

type Day4Solution1 struct {}

func (Day4Solution1) Solve(path string) {
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")

	fmt.Println(findXmas(lines))
}