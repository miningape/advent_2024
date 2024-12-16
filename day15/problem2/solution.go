package day15_problem2

import (
	"advent2024/day15"
	"advent2024/util"
	"fmt"
	"time"
)

func widenWarehouse(warehouse *util.Grid[day15.Cell]) {
	doubled := make(util.Grid[day15.Cell], 0, len(*warehouse) * 2)

	for _, line := range *warehouse {
		doubledLine := make([]day15.Cell, 0, len(line) * 2)

		for _, cell := range line {
			switch cell {
			case day15.Wall:
				doubledLine = append(doubledLine, day15.Wall, day15.Wall)
			case day15.Empty:
				doubledLine = append(doubledLine, day15.Empty, day15.Empty)
			case day15.Box:
				doubledLine = append(doubledLine, day15.BoxLeft, day15.BoxRight)
			default:
				panic(fmt.Sprint("Cannot recognize: ", cell.ToString()))
			}
		}

		doubled = append(doubled, doubledLine)
	}

	*warehouse = doubled
}

func moveRobotToWidenedWarehouse(robot *util.Vector) {
	robot.X *= 2
}

func shift(location util.Vector, direction util.Vector, warehouse *util.Grid[day15.Cell]) {
	current, _ := warehouse.At(location)
	warehouse.Set(location, day15.Empty)
	warehouse.Set(location.Add(direction), current)
}

func moveBoxHorisontally(boxLocation util.Vector, direction util.Vector, warehouse *util.Grid[day15.Cell]) bool {
	next := boxLocation.Add(direction)
	for {
		cell, found := warehouse.At(next)
		if !found {
			return false
		}

		switch cell {
		case day15.Wall:
			return false

		case day15.BoxLeft:
			fallthrough
		case day15.BoxRight:
			next = next.Add(direction)

		case day15.Empty:
			for next.VectorMul(direction).ManhattanOrigin() != boxLocation.VectorMul(direction).ManhattanOrigin() {
				next = next.Sub(direction)
				shift(next, direction, warehouse)
			}

			return true

		default:
			panic(fmt.Sprint("Unrecognized symbol ", cell.ToString()))
		}
	}
}

func findOther(boxLocation util.Vector, warehouse *util.Grid[day15.Cell]) util.Vector {
	boxPart := (*warehouse)[boxLocation.Y][boxLocation.X]

	if boxPart == day15.BoxLeft {
		return boxLocation.Add(util.Vector{X: 1, Y: 0})
	}

	if boxPart == day15.BoxRight {
		return boxLocation.Add(util.Vector{X: -1, Y: 0})
	}
 
	fmt.Println(boxPart.ToString(), boxLocation)
	panic("Not a box")
}

func canMoveBoxVertically(location util.Vector, warehouse *util.Grid[day15.Cell]) (bool, bool) {
	cell, found := warehouse.At(location)
	if !found {
		return false, false
	}

	switch cell {
	case day15.Wall:
		return false, false
	
	case day15.BoxRight:
		fallthrough
	case day15.BoxLeft:
		return false, true
	
	case day15.Empty:
		return true, true
	
	default:
		panic(fmt.Sprint("Unrecognized symbol ", cell.ToString()))
	}
}

func moveBoxVertically(boxLocation util.Vector, direction util.Vector, warehouse *util.Grid[day15.Cell]) bool {
	boxOtherLocation := findOther(boxLocation, warehouse)

	canMove, canMoveIfNextMoves := canMoveBoxVertically(boxLocation.Add(direction), warehouse)
	if !canMove && !canMoveIfNextMoves {
		return false
	}

	snapshot := warehouse.DeepClone()
	if !canMove {
		moved := moveBoxVertically(boxLocation.Add(direction), direction, warehouse)

		if !moved {
			return false
		}
	}

	canMoveOther, canMoveOtherIfNextMoves := canMoveBoxVertically(boxOtherLocation.Add(direction), warehouse)
	if !canMoveOther && !canMoveOtherIfNextMoves {
		*warehouse = snapshot
		return false
	}

	if !canMoveOther {
		moved := moveBoxVertically(boxOtherLocation.Add(direction), direction, warehouse)

		if !moved {
			*warehouse = snapshot
			return false
		}
	}

	shift(boxLocation, direction, warehouse)
	shift(boxOtherLocation, direction, warehouse)

	return true
}

func moveBox(boxLocation util.Vector, direction util.Vector, warehouse *util.Grid[day15.Cell]) bool {
	if direction.Y == 0 {
		return moveBoxHorisontally(boxLocation, direction, warehouse)
	}

	if direction.X == 0 {
		return moveBoxVertically(boxLocation, direction, warehouse)
	}

	panic("Cannot move box diagonally")
}

func move(instruction util.Vector, robot *util.Vector, warehouse *util.Grid[day15.Cell]) {
	next := robot.Add(instruction)

	cell, found := warehouse.At(next)
	if !found {
		panic("Tried to leave the warehouse")
	}

	switch cell {
	case day15.Wall:
	case day15.Empty:
		*robot = next

	case day15.BoxLeft:
		fallthrough
	case day15.BoxRight:
		moved := moveBox(next, instruction, warehouse)
		if moved {
			*robot = next
		}

	default:
		panic(fmt.Sprint("Could not recognize cell: ", cell.ToString()))
	}
}

func findAllBoxes(grid *util.Grid[day15.Cell]) []util.Vector {
	boxes := make([]util.Vector, 0)

	for y, line := range *grid {
		for x, c := range line {
			if c == day15.BoxLeft {
				boxes = append(boxes, util.Vector{X: x, Y: y})
			}
		}
	}

	return boxes
}

type Day15Solution2 struct {}

func (Day15Solution2) Solve(path string) {
	defer util.MeasureRuntime(time.Now())
	file := util.ReadFile(path)
	robot, warehouse, instructions := day15.ParseInput(file)

	widenWarehouse(&warehouse)
	moveRobotToWidenedWarehouse(&robot)

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