package day8_problem2

import (
	"advent2024/day08"
	"advent2024/util"
	"fmt"
	"math"
	"strings"
)

type counter map[int]int;

func (c counter) has(n int) bool {
	v, found := c[n]

	return found && v > 0
}

func (c counter) add(n int) {
	v, found := c[n]

	if !found {
		c[n] = 1
		return
	}

	c[n] = v + 1
}

func (c counter) remove(n int) {
	if !c.has(n) {
		panic("Tried to remove number that doesn't exist from counter")
	}

	c[n] = c[n] - 1
}

func (c counter) product() int {
	sum := 1

	for num, count := range c {
		for i := 0; i < count; i++ {
			sum *= num
		}
	}

	return sum
}

func findPrimeFactors(num int) counter {
	factors := make(counter)

	if num == 0 {
		factors.add(0)
		return factors
	}

	if num < 0 {
		num *= -1
		factors.add(-1)
	}

	for num % 2 == 0 {
		num /= 2
		factors.add(2)
	}

	max := int(math.Ceil(math.Sqrt(float64(num)))) 
	for i := 3; i <= max; i += 2 {
		for num % i == 0 {
			num /= i
			factors.add(i)
		}
	}

	if num > 2 {
		factors.add(num)
	}

	return factors
}

/*
	Turns out the input doesn't allow for any antinodes between each antenna - this code accounts for that :P 
*/
func findSmallestVectorAlong(direction util.Vector) util.Vector {
	if direction.X == 0 {
		if direction.Y == 0 {
			panic("Cannot find direction in vector with 0, 0")
		}

		return util.Vector{
			X: 0,
			Y: 1,
		}
	}

	if direction.Y == 0 {
		return util.Vector{
			X: 1,
			Y: 0,
		}
	}

	primes_x := findPrimeFactors(direction.X)
	primes_y := findPrimeFactors(direction.Y)

	for num := range primes_x {
		for primes_y.has(num) && primes_x.has(num)  {
			primes_x.remove(num)
			primes_y.remove(num)
		}
	}

	return util.Vector{
		X: primes_x.product(),
		Y: primes_y.product(),
	}
}

func findAntiNodes(lines []string, antennas map[rune][]util.Vector) util.Set[util.Vector] {
	antiNodes := util.SetOf[util.Vector]()

	for _, antennas := range antennas {
		for i, source := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				other := antennas[j]

				direction := other.Sub(source)
				smallest := findSmallestVectorAlong(direction)

				next := source
				for day08.IsInside(lines, next) {
					antiNodes.Add(next)
					next = next.Add(smallest)
				}

				next = source.Sub(smallest)
				for day08.IsInside(lines, next) {
					antiNodes.Add(next)
					next = next.Sub(smallest)
				}
			}
		}
	}

	return antiNodes
}

type Day8Solution2 struct {}

func printMap(lines []string, antiNodes util.Set[util.Vector]) {
	for y, line := range lines {
		for x, c := range line {
			if c == '.' && antiNodes.Contains(util.Vector{ X: x, Y: y }) {
				fmt.Print("#")
			} else {
				fmt.Print(string(c))
			}
		}

		fmt.Println()
	}
}

func (Day8Solution2) Solve(path string) {
	file := util.ReadFile(path)
	lines := strings.Split(file, "\n")

	antennas := day08.FindAntennas(lines)
	antiNodes := findAntiNodes(lines, antennas)

	// printMap(lines, antiNodes)
	fmt.Println("Answer:", len(antiNodes))
}