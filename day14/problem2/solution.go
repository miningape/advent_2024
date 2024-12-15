package day14_problem2

import (
	"advent2024/day14"
	"advent2024/util"
	"fmt"
	"strconv"
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

func print(robots util.Counter[util.Vector], bounds util.Vector) {
	for y := 0; y < bounds.Y; y++ {
		for x := 0; x < bounds.X; x++ {
			position := util.Vector{X: x, Y: y}
			if robots.Has(position) {
				fmt.Print(strconv.Itoa(robots[position]))
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func hasStraightLine(robots util.Counter[util.Vector], length int) bool {
	for robot := range robots {
		for i := 1; i <= length; i++ {
			if !robots.Has(robot.Add(util.Vector{X: 0, Y: i})) {
				break
			}  
			
			if i == length {
				return true
			}
		}
	}

	return false
}

type Day14Solution2 struct {}

func (Day14Solution2) Solve(path string) {
	file := util.ReadFile(path)
	robots := day14.ParseInput(file)

	bounds := util.Vector{X: 101, Y: 103}

	max := bounds.X * bounds.Y // Min moves for minimally moving (1, 1) robot to visit every node - we repeat positions after this
	for i := 0; i < max; i++ {
		positions := robotsAfter(robots, i, bounds)
		
		if hasStraightLine(positions, 10) {
			print(positions, bounds)
			fmt.Println("==>", i)
		}
	}
}