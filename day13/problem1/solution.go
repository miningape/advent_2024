package day13_problem1

import (
	"advent2024/day13"
	"advent2024/util"
	"fmt"
)

func dynamicTokensToWin(clawMachine day13.ClawMachine) (int, bool) {
	tokens := make(map[util.Vector]int)
	tokens[util.Vector{X: 0, Y: 0}] = 0
	tokens[clawMachine.Controls.A] = 3
	tokens[clawMachine.Controls.B] = 1

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			position := clawMachine.Controls.A.Mul(a).Add(clawMachine.Controls.B.Mul(b))
			_, found := tokens[position]
			if found {
				continue
			}

			costA, foundA := tokens[position.Sub(clawMachine.Controls.A)]
			costA += 3
			costB, foundB := tokens[position.Sub(clawMachine.Controls.B)]
			costB += 1

			if !foundA && !foundB {
				continue
			}

			if foundA && foundB {
				if costA < costB {
					tokens[position] = costA 
				} else {
					tokens[position] = costB
				}
			}

			if foundA && !foundB {
				tokens[position] = costA
			}

			if !foundA && foundB {
				tokens[position] = costB
			}
		}
	}

	cost, found := tokens[clawMachine.Prize]
	return cost, found
}

func pressA(machine day13.ClawMachine) day13.ClawMachine {
	return day13.ClawMachine{
		Controls: machine.Controls,
		Prize: machine.Prize.Sub(machine.Controls.A),
	}
}

func pressB(machine day13.ClawMachine) day13.ClawMachine {
	return day13.ClawMachine{
		Controls: machine.Controls,
		Prize: machine.Prize.Sub(machine.Controls.B),
	}
}

func recursiveTokensToWin(clawMachine day13.ClawMachine) (int, bool) {
	if clawMachine.Prize == (util.Vector{}) {
		return 0, true
	}

	if clawMachine.Prize.X < 0 || clawMachine.Prize.Y < 0 {
		return 0, false
	}

	tokensA, winA := recursiveTokensToWin(pressA(clawMachine))
	tokensA += 3
	tokensB, winB := recursiveTokensToWin(pressB(clawMachine))
	tokensB += 1

	if winA && winB {
		if tokensA < tokensB {
			return tokensA, true
		}
	
		return tokensB, true
	}

	if winA && !winB {
		return tokensA, true
	}

	if !winA && winB {
		return tokensB, true
	}

	
	return 0, false
}

type Day13Solution1 struct {}

func (Day13Solution1) Solve(path string) {
	file := util.ReadFile(path)
	clawMachines := day13.ParseInput(file)

	total := 0
	for _, clawMachine := range clawMachines {
		tokens, possible := dynamicTokensToWin(clawMachine)

		if possible {
			total += tokens
		}
	}

	fmt.Println(total)
}