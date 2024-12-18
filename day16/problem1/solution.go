package day16_problem1

import (
	"advent2024/util"
	"fmt"
	"strings"
)

type Moment struct {
	Position util.Vector
	Direction util.Vector
}

func ParseInput(file string) (util.Grid[bool], util.Vector, util.Vector) {
	lines := strings.Split(file, "\n")
	start := util.Vector{X: -1, Y: -1}
	end := util.Vector{X: -1, Y: -1}

	maze := util.GridOfString(lines, func(r rune, v util.Vector) bool {
		switch r {
		case '#':
			return true
		case '.':
			return false

		case 'E':
			end = v
			return false
		case 'S':
			start = v
			return false
		
		default:
			panic(fmt.Sprint("Cannot recognise rune: ", string(r)))
		}
	})

	return maze, start, end 
}

func toString(grid util.Grid[bool], start, end util.Vector, path []util.Vector) string {
	special := map[util.Vector]string{
		end: "E",
	}
	for _, step := range path {
		special[step] = "*"
	}
	special[start] = "S"

	s, err := util.GridToString(grid, map[bool]string{
		true: "#",
		false: ".",
	}, special)

	if err != nil {
		panic(err)
	}

	return s
}

func shortestPath(maze util.Grid[bool], start, end util.Vector) ([]util.Vector, int) {
	type CostVector struct {
		vec util.Vector
		cost int
	}

	pq := util.PriorityQueueOf(Moment{start, util.EAST}, 0)
	parent := make(map[util.Vector]CostVector)

	for !pq.IsEmpty() {
		moment, cost := pq.PullPriority()

		forward := moment.Position.Add(moment.Direction)
		isWall, found := maze.At(forward)
		if found && !isWall {
			_, found := parent[forward]
			if !found {
				parent[forward] = CostVector{moment.Position, 1}
				next := Moment{forward, moment.Direction}
				pq.Insert(next, cost + 1)
			}
		}

		right := moment.Direction.RotateOrigin90()
		forward = moment.Position.Add(right)
		isWall, found = maze.At(forward)
		if found && !isWall {
			_, found := parent[forward]
			if !found {
				parent[forward] = CostVector{moment.Position, 1001}
				next := Moment{forward, right}
				pq.Insert(next, cost + 1001)
			}
		}

		left := moment.Direction.RotateOrigin90().Opposite()
		forward = moment.Position.Add(left)
		isWall, found = maze.At(forward)
		if found && !isWall {
			_, found := parent[forward]
			if !found {
				parent[forward] = CostVector{moment.Position, 1001}
				next := Moment{forward, left}
				pq.Insert(next, cost + 1001)
			}
		}
	}

	total := 0
	path := make([]util.Vector, 0)
	cv, found := parent[end]
	for found {
		path = append(path, cv.vec)
		total += cv.cost
		cv, found = parent[cv.vec]
	}

	return path, total
}

type Day16Solution1 struct {}

func (Day16Solution1) Solve(path string) {
	file := util.ReadFile(path)

	maze, start, end := ParseInput(file)
	p, l := shortestPath(maze, start, end)

	fmt.Println(toString(maze, start, end, p))
	fmt.Println(len(p))
	fmt.Println(l)
}