package day17_problem2

import (
	"advent2024/day17"
	"advent2024/util"
	"fmt"
)

func quine(computer day17.Computer, index int, register int) (int, bool) {
	if index < 0 {
		return register >> 3, true
	}

	// Trying all 3-bit numbers
	for i := 0; i < 8; i++ {
		computer.Registers.A = register + i

		for len(computer.Output) == 0 {
			computer = computer.Cycle()
		}

		computer.InstructionPointer = 0
		computer.Registers.B = 0
		computer.Registers.C = 0
		
		value := computer.Output[0]
		computer.Output = make([]int, 0)
		
		if value == computer.Program[index] {
			quineGuess := (register + i) << 3
			quine, found := quine(computer, index - 1, quineGuess)
			if found {
				return quine, true
			}
		}
	}

	return 0, false
}

type Day17Solution2 struct{}

func (Day17Solution2) Solve(path string) {
	file := util.ReadFile(path)
	computer := day17.ParseInput(file)

	quine, found := quine(computer, len(computer.Program) - 1, 0)
	if !found {
		panic("Could not find a quine")
	}
	
	fmt.Println(quine)
}
