package day15

import (
	"advent2024/util"
	"fmt"
	"strings"
)

type Cell int8;
const (
	Empty Cell = iota
	Wall
	Box
	BoxLeft
	BoxRight
)

func (c Cell) ToString() string {
	switch c {
	case Empty:
		return "."
	case Wall:
		return "#"
	case Box:
		return "O"
	case BoxLeft:
		return "["
	case BoxRight:
		return "]"
	}

	panic("Unrecognized")
}

type Matcher struct {
	*util.Matcher
}

func symbolToVector(symbol rune) util.Vector {
	switch symbol {
	case '<':
		return util.Vector{X: -1, Y: 0}
	case '^':
		return util.Vector{X: 0, Y: -1}
	case '>':
		return util.Vector{X: 1, Y: 0}
	case 'v':
		return util.Vector{X: 0, Y: 1}
	}

	panic(fmt.Sprint("Could not recognise: ", string(symbol)))
}

func (matcher *Matcher) readInstructions() []util.Vector {
	instructions := make([]util.Vector, 0, )

	for !matcher.IsAtEnd() {
		matcher.Consume('\n')

		symbol, found := matcher.Next()
		if found {
			instruction := symbolToVector(symbol)
			instructions = append(instructions, instruction)
		}
	}

	return instructions
}


func ParseInput(file string) (util.Vector, util.Grid[Cell], []util.Vector) {
	f := strings.Split(file, "\n\n")
	robot := util.Vector{X: -1, Y: -1}

	warehouse := util.GridOfString(strings.Split(f[0], "\n"), func(r rune, location util.Vector) Cell {
		switch r {
		case '.':
			return Empty
		case '#':
			return Wall
		case 'O':
			return Box
		case '@':
			robot = location
			return Empty
		default:
			panic(fmt.Sprint("Cannot recognise stymbol ", string(r)))
		}
	})

	if robot.X == -1 && robot.Y == -1 {
		panic("Could not find robot")
	}
	
	matcher := Matcher{&util.Matcher{Source: f[1], Index: 0}}
	instructions := matcher.readInstructions()

	return robot, warehouse, instructions
}

func ToString(grid util.Grid[Cell], robot util.Vector) string {
	g := ""

	for y, line := range grid {
		for x, c := range line {
			if x == robot.X && y == robot.Y {
				g += "@"
			} else {
				g += c.ToString()
			}
		}

		g += "\n"
	}

	return g
}