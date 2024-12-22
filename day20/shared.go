package day20

import (
	"advent2024/util"
	"fmt"
	"strings"
)

func ParseInput(file string) (util.Grid[bool], util.Vector, util.Vector, util.Set[util.Vector]) {
	lines := strings.Split(file, "\n")
	path := util.SetOf[util.Vector]()

	var start, end util.Vector
	grid := util.GridOfString(lines, func(r rune, v util.Vector) bool {
		switch r {
		case '#':
			return true
		case '.':
			path.Add(v)
			return false
		case 'S':
			start = v
			path.Add(v)
			return false
		case 'E':
			end = v
			path.Add(v)
			return false
		}

		panic(fmt.Sprint("Cannot recognize ", string(r), v))
	})

	return grid, start, end, path
}

func MazeToString(maze util.Grid[bool], path []util.Vector) string {
	positional := make(map[util.Vector]string)
	for _, step := range path {
		positional[step] = "O"
	}

	s, err := util.GridToString(maze, map[bool]string{
		true: "#",
		false: ".",
	}, positional)

	if err != nil {
		panic(err)
	}

	return s
}
