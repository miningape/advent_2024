package day13_problem2

import (
	"advent2024/day13"
	"advent2024/util"
	"fmt"
)

func linearTokensToWin(clawMachine day13.ClawMachine) (int, bool) {
	equations := util.LinearEquations(clawMachine.Controls.A, clawMachine.Controls.B)
	presses, solved := equations.SolveFor(clawMachine.Prize)

	if !solved {
		return 0, false
	}

	tokensA := presses.X * 3
	tokensB := presses.Y

	return tokensA + tokensB, true
}

type Day13Solution2 struct {}

func (Day13Solution2) Solve(path string) {
	file := util.ReadFile(path)
	clawMachines := day13.ParseInput(file)

	total := 0
	for _, clawMachine := range clawMachines {
		tokens, possible := linearTokensToWin(clawMachine)
		
		if possible {
			total += tokens
		}
	}

	fmt.Println(total)
}