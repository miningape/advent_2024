package util

type Grid[T any] map[Vector]T

func GridOfRaw(m []string) Grid[rune] {
	grid := make(Grid[rune])

	for y, line := range m {
		for x, c := range line {
			grid[Vector{x, y}] = c
		}
	}

	return grid
}