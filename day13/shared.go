package day13

import "advent2024/util"

type ClawMachine struct {
	Controls struct {
		A util.Vector
		B util.Vector
	}
	Prize util.Vector
}

type Matcher struct {
	*util.Matcher
}

func (matcher *Matcher) matchButton(button *util.Vector, letter string) {
	matcher.Match("Button " + letter + ": X+")
	button.X = matcher.Match_int()
	matcher.Match(", Y+")
	button.Y = matcher.Match_int()
	matcher.Match("\n")
}

func (matcher *Matcher) matchPrize(clawMachine *ClawMachine) {
	matcher.Match("Prize: X=")
	clawMachine.Prize.X = matcher.Match_int()
	matcher.Match(", Y=")
	clawMachine.Prize.Y = matcher.Match_int()
}

func (matcher *Matcher) matchClawMachine() ClawMachine {
	clawMachine := ClawMachine{}

	matcher.matchButton(&clawMachine.Controls.A, "A")
	matcher.matchButton(&clawMachine.Controls.B, "B")
	matcher.matchPrize(&clawMachine)

	return clawMachine
}

func (matcher *Matcher) matchClawMachines() []ClawMachine {
	machines := make([]ClawMachine, 0)

	for !matcher.IsAtEnd() {
		machine := matcher.matchClawMachine()
		machine.Prize = machine.Prize.Add(util.Vector{X: 1, Y: 1}.Mul(10000000000000))
		machines = append(machines, machine)

		if !matcher.IsAtEnd() {
			matcher.Match("\n\n")
		}
	}

	return machines
}

func ParseInput(file string) []ClawMachine {
	matcher := Matcher{&util.Matcher{Source: file, Index: 0}}
	machines := matcher.matchClawMachines()

	return machines
}