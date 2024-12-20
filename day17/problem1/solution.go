package day17_problem1

import (
	"advent2024/util"
	"fmt"
	"strconv"
)

type Matcher struct {
	*util.Matcher
}

type Instruction struct {
	Opcode int
	Operand int
}

func (instruction Instruction) execute(computer Computer) Computer {
	switch instruction.Opcode {
	case 0: // adv
		computer.InstructionPointer += 2
		computer.Registers.A = computer.Registers.A / util.Pow(2, computer.combo(instruction.Operand))
	
	case 1: // bxl
		computer.InstructionPointer += 2
		computer.Registers.B = computer.Registers.B ^ instruction.Operand
	
	case 2: // bst
		computer.InstructionPointer += 2
		computer.Registers.B = computer.combo(instruction.Operand) % 8
	
	case 3: // jnz
		if computer.Registers.A == 0 {
			computer.InstructionPointer += 2
		} else {
			computer.InstructionPointer = instruction.Operand
		}
	
	case 4: // bxc
		computer.InstructionPointer += 2
		computer.Registers.B = computer.Registers.B ^ computer.Registers.C

	case 5: // out
		computer.InstructionPointer += 2
		value := strconv.Itoa(computer.combo(instruction.Operand) % 8)

		if len(computer.Output) != 0 {
			value = "," + value
		}

		computer.Output += value

	case 6: // bdv
		computer.InstructionPointer += 2
		computer.Registers.B = computer.Registers.A / util.Pow(2, computer.combo(instruction.Operand))

	case 7: // cdv
		computer.InstructionPointer += 2
		computer.Registers.C = computer.Registers.A / util.Pow(2, computer.combo(instruction.Operand))

	default:
		panic(fmt.Sprint("Opcode out of range (> 7): ", instruction.Opcode))
	}

	return computer
}

type Computer struct {
	InstructionPointer int
	Program []int
	Registers struct {
		A int
		B int
		C int
	}
	Output string
}

func (computer Computer) combo(operand int) int {
	switch operand {
	case 0: fallthrough
	case 1: fallthrough
	case 2: fallthrough
	case 3:
		return operand
	case 4:
		return computer.Registers.A
	case 5:
		return computer.Registers.B
	case 6:
		return computer.Registers.C
	case 7:
		panic("Reserved operand (7)")
	default:
		panic(fmt.Sprint("Combo-operand out of range (> 7): ", operand))
	}
}

func (computer Computer) isHalted() bool {
	return computer.InstructionPointer >= len(computer.Program) - 1
} 

func (computer Computer) currentInstruction() Instruction {
	if computer.InstructionPointer < 0 {
		panic("Instruction pointer is before the program")
	}

	if computer.isHalted() {
		panic("Cannot get the next instruction for a halted computer")
	}

	return Instruction{
		Opcode: computer.Program[computer.InstructionPointer],
		Operand: computer.Program[computer.InstructionPointer + 1],
	}
}

func (computer Computer) next() Computer {
	instruction := computer.currentInstruction()
	return instruction.execute(computer)
}

func (matcher Matcher) parseInput() Computer {
	computer := Computer{
		InstructionPointer: 0, 
		Program: make([]int, 0), 
		Registers: struct{A int; B int; C int}{ A: 0, B: 0, C: 0 },
		Output: "",
	}

	matcher.Match("Register A: ")
	computer.Registers.A = matcher.Match_int()
	matcher.Match("\n")

	matcher.Match("Register B: ")
	computer.Registers.B = matcher.Match_int()
	matcher.Match("\n")

	matcher.Match("Register C: ")
	computer.Registers.C = matcher.Match_int()
	matcher.Match("\n")

	matcher.Match("\nProgram: ")
	for !matcher.IsAtEnd() {
		instruction := matcher.Match_int()
		matcher.Consume(',')

		computer.Program = append(computer.Program, instruction)
	}

	return computer
}

type Day17Solution1 struct{}

func (Day17Solution1) Solve(path string) {
	file := util.ReadFile(path)

	matcher := Matcher{&util.Matcher{Source: file, Index: 0}}
	computer := matcher.parseInput()

	for !computer.isHalted() {
		computer = computer.next()
	}

	fmt.Println(computer.Output)
}