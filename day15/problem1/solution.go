package day15_problem1

import (
	"advent2024/day15"
	"advent2024/util"
	"fmt"
)

func move(direction util.Vector, robot *util.Vector, grid *util.Grid[day15.Cell]) {
	next := robot.Add(direction)

	cell, found := grid.At(next)
	if !found {
		panic("Tried to leave the warehouse")
	}

	if cell == day15.Wall {
		return
	}

	if cell == day15.Empty {
		*robot = next
		return
	}

	if cell == day15.Box {
		search := next.Add(direction)
		for {
			searched, found := grid.At(search)
			if !found {
				return
			}

			if searched == day15.Wall {
				return
			}

			if searched == day15.Empty {
				break
			}

			search = search.Add(direction)
		}

		grid.Set(search, day15.Box)
		grid.Set(next, day15.Empty)
		*robot = next

		return
	}

	panic(fmt.Sprint("move - Cannot recognize ", cell))
}

func findAllBoxes(grid *util.Grid[day15.Cell]) []util.Vector {
	boxes := make([]util.Vector, 0)

	for y, line := range *grid {
		for x, c := range line {
			if c == day15.Box {
				boxes = append(boxes, util.Vector{X: x, Y: y})
			}
		}
	}

	return boxes
}

type Day15Solution1 struct {}

func (Day15Solution1) Solve(path string) {
	file := util.ReadFile(path)
	robot, warehouse, instructions := day15.ParseInput(file)

	for _, instruction := range instructions {
		move(instruction, &robot, &warehouse)
	}

	sum := 0
	boxes := findAllBoxes(&warehouse)
	for _, box := range boxes {
		sum += box.X + (box.Y * 100)
	}

	fmt.Println(sum)
}