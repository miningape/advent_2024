package day16_problem2

import (
	"advent2024/util"
	"fmt"
	"math"
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

func isEmpty(location util.Vector, maze util.Grid[bool]) bool {
	isWall, found := maze.At(location)
	return found && !isWall
}

type CostVectorDirection struct {
	cost int
	vec util.Vector
	dir util.Vector
}

func children2(location, direction util.Vector, maze util.Grid[bool]) []CostVectorDirection {
	child := make([]CostVectorDirection, 0)

	forward := location.Add(direction)
	if isEmpty(forward, maze) {
		child = append(child, CostVectorDirection{vec: forward, dir: direction, cost: 1})
	}

	left := location.Add(direction.RotateOrigin90())
	if isEmpty(left, maze) {
		child = append(child, CostVectorDirection{vec: location, dir: direction.RotateOrigin90(), cost: 1000})
	}

	right := location.Add(direction.RotateOrigin90().Opposite())
	if isEmpty(right, maze) {
		child = append(child, CostVectorDirection{vec: location, dir: direction.RotateOrigin90().Opposite(), cost: 1000})
	}

	return child
}

type CostSet struct{
	path util.Set[util.Vector]
	cost int
}

func findAllPossiblePathsWithScore(start, direction util.Vector, maze util.Grid[bool]) map[util.Vector]map[util.Vector]CostSet {
	cache := make(map[util.Vector]map[util.Vector]CostSet)

	cache[start] = make(map[util.Vector]CostSet)
	cache[start][direction] = CostSet{path: util.SetOf(start), cost: 0}

	pq := util.PriorityQueueOf(start, 0)
	for !pq.IsEmpty() {
		current := pq.PullValue()

		previous, found := cache[current]

		if !found {
			panic("Cannot perform next movement if previous is not available")
		}

		for direction, pathCost := range previous {
			for _, child := range children2(current, direction, maze) {
				_, found := cache[child.vec]
				if !found {
					cache[child.vec] = make(map[util.Vector]CostSet)
				}

				cached, found := cache[child.vec][child.dir]
				cost := child.cost + pathCost.cost
				if !found {
					pq.Insert(child.vec, cost)
					cache[child.vec][child.dir] = CostSet{pathCost.path.Union(util.SetOf(child.vec)), cost}
				} else if cost == cached.cost {
					cache[child.vec][child.dir] = CostSet{cached.path.Union(pathCost.path.Union(util.SetOf(child.vec))), cost}
				}
			}
		}
	}

	return cache
}

func findShortestPathCost(paths map[util.Vector]map[util.Vector]CostSet, end util.Vector) int {
	shortest := math.MaxInt

	for _, cardinal := range util.Cardinals() {
		if paths[end][cardinal].cost < shortest {
			shortest = paths[end][cardinal].cost
		}
	}

	if shortest == math.MaxInt {
		panic("Could not find any paths to the end")
	}

	return shortest
}

func cellsAlongShortestPaths(paths map[util.Vector]map[util.Vector]CostSet, end util.Vector, cost int) util.Set[util.Vector] {
	cells := util.SetOf[util.Vector]()
	for _, cardinal := range util.Cardinals() {
		if cost == paths[end][cardinal].cost {
			for step := range paths[end][cardinal].path {
				cells.Add(step)
			}
		}
	}

	return cells
}

type Day16Solution2 struct {}

func (Day16Solution2) Solve(path string) {
	file := util.ReadFile(path)

	maze, start, end := ParseInput(file)
	paths := findAllPossiblePathsWithScore(start, util.EAST, maze)
	score := findShortestPathCost(paths, end)
	cells := cellsAlongShortestPaths(paths, end, score)

	fmt.Println(len(cells), score)
	fmt.Println(toString(maze, start, end, cells.Slice()))
}