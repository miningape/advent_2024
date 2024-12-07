package day6_problem2

import (
	"advent2024/util"
	"fmt"
	"maps"
	"strings"
	"time"
)

type moment struct {
	position util.Vector
	direction util.Vector
}

func findStartingPostion(lines []string) (util.Vector, util.Vector) {
	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				return util.Vector{X: x, Y: y}, util.Vector{X: 0, Y: -1}
			}
		}
	}

	panic("Could not find \"^\" (starting position) in map.")
}

func isInside(lines []string, position util.Vector) bool {
	if position.Y < 0 || position.Y > len(lines)-1 {
		return false
	}

	if position.X < 0 || position.X > len(lines[position.Y])-1 {
		return false
	}

	return true
}

func nextPosition(lines []string, position util.Vector, direction util.Vector) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && rune(lines[next.Y][next.X]) == '#' {
		return  position, direction.RotateOrigin90().Opposite()
	}

	return next, direction
}

func nextPositionWithObstacle(lines []string, position util.Vector, direction util.Vector, obstacle util.Vector) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && (next == obstacle || rune(lines[next.Y][next.X]) == '#') {
		return  position, direction.RotateOrigin90().Opposite()
	}

	return next, direction
}

func isLoop(lines []string, path util.Set[moment], position util.Vector, direction util.Vector, obstacle util.Vector) bool {
	seen := maps.Clone(path)

	for isInside(lines, position) {
		current := moment{ position, direction }
		if seen.Contains(current) {
			return true
		}
		
		seen.Add(current)
		position, direction = nextPositionWithObstacle(lines, position, direction, obstacle)
	}

	return false
}

func walkUntilLeaves(lines []string, position util.Vector, direction util.Vector) util.Set[util.Vector] {
	path := util.SetOf[moment]()
	stepped_on := util.SetOf[util.Vector]()
	loopObstacles := util.SetOf[util.Vector]()

	for isInside(lines, position) {
		path.Add(moment{ position, direction })
		stepped_on.Add(position)
		
		in_front := position.Add(direction)
		if  !stepped_on.Contains(in_front) && isLoop(lines, path, position,  direction.RotateOrigin90().Opposite(), in_front) {
			loopObstacles.Add(in_front)
		}

		position, direction = nextPosition(lines, position, direction)
	}

	return loopObstacles
}

func writeOutput(lines []string, obstacles util.Set[util.Vector]) {
	for y, line := range lines {
		for x, c := range line {
			if obstacles.Contains(util.Vector{ X: x, Y: y }) {
				fmt.Print("O")
			} else {
				fmt.Print(string(c))
			}
		}

		fmt.Print("\n")
	}
}

type Day5Solution2 struct{}

func (Day5Solution2) Solve(path string) {
	start := time.Now()
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")

	position, direction := findStartingPostion(lines)
	obstacles := walkUntilLeaves(lines, position, direction)

	fmt.Println(len(obstacles))
	elapsed := time.Since(start)
	fmt.Println("Took:", elapsed)
}
