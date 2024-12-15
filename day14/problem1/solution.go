package day14_problem1

import (
	"advent2024/day14"
	"advent2024/util"
	"fmt"
)

func robotsAfter(robots []day14.Robot, seconds int, bounds util.Vector) util.Counter[util.Vector] {
	positions := util.CounterOf[util.Vector]()

	for _, robot := range robots {
		next := robot.Position.Add(robot.Velocity.Mul(seconds))
		actual := util.Vector{
			X: next.X % bounds.X,
			Y: next.Y % bounds.Y,
		}

		if actual.X < 0 {
			actual.X = bounds.X + actual.X
		}

		if actual.Y < 0 {
			actual.Y = bounds.Y + actual.Y
		}

		positions.Add(actual)
	}

	return positions
}

func robotsInQuadrant(x, y int, robots util.Counter[util.Vector], bounds util.Vector) int {
	min := util.Vector{
		X: x * ((bounds.X / 2) + (bounds.X % 2)), 
		Y: y * ((bounds.Y / 2) + (bounds.Y % 2)),
	}
	max := util.Vector{
		X: (bounds.X / 2) + min.X, 
		Y: (bounds.Y / 2) + min.Y,
	}

	count := 0
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			position := util.Vector{
				X: x,
				Y: y,
			}

			if robots.Has(position) {
				count += robots[position]
			} else {
			}
		}
	}

	return count
}

type Day14Solution1 struct {}

func (Day14Solution1) Solve(path string) {
	file := util.ReadFile(path)
	robots := day14.ParseInput(file)

	bounds := util.Vector{X: 101, Y: 103}
	positions := robotsAfter(robots, 100, bounds)
	
	northeast := robotsInQuadrant(0, 0, positions, bounds)
	northwest := robotsInQuadrant(0, 1, positions, bounds)
	southeast := robotsInQuadrant(1, 0, positions, bounds)
	southwest := robotsInQuadrant(1, 1, positions, bounds)

	fmt.Println(northeast * northwest * southeast * southwest)
}