package day18

import (
	"advent2024/util"
	"strconv"
	"strings"
)

func ParseInput(file string) []util.Vector {
	lines := strings.Split(file, "\n")
	corrupted := make([]util.Vector, len(lines))

	for i, line := range lines {
		l := strings.Split(line, ",")
		
		x, err := strconv.Atoi(l[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}

		corrupted[i] = util.Vector{
			X: x,
			Y: y,
		}
	}

	return corrupted
}

func ToString(memory util.Grid[bool], path []util.Vector) string {
	p := make(map[util.Vector]string)
	for _, step := range path {
		p[step] = "O"
	}

	 s, err := util.GridToString(memory, map[bool]string {
		true: "#",
		false: ".",
	}, p)

	if err != nil {
		panic(err)
	}

	return s
}

func ApplyCorruption(corruptions []util.Vector, grid util.Grid[bool]) {
	for _, corruption := range corruptions {
		grid.Set(corruption, true)
	}
}