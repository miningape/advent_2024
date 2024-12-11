package main

import (
	day1_problem1 "advent2024/day01/problem1"
	day1_problem2 "advent2024/day01/problem2"
	day2_problem1 "advent2024/day02/problem1"
	day2_problem2 "advent2024/day02/problem2"
	day3_problem1 "advent2024/day03/problem1"
	day3_problem2 "advent2024/day03/problem2"
	day4_problem1 "advent2024/day04/problem1"
	day4_problem2 "advent2024/day04/problem2"
	day5_problem1 "advent2024/day05/problem1"
	day5_problem2 "advent2024/day05/problem2"
	day6_problem1 "advent2024/day06/problem1"
	day6_problem2 "advent2024/day06/problem2"
	day7_problem1 "advent2024/day07/problem1"
	day7_problem2 "advent2024/day07/problem2"
	day8_problem1 "advent2024/day08/problem1"
	day8_problem2 "advent2024/day08/problem2"
	day9_problem1 "advent2024/day09/problem1"
	day9_problem2 "advent2024/day09/problem2"
	day10_problem1 "advent2024/day10/problem1"
	day10_problem2 "advent2024/day10/problem2"
	"advent2024/util"
	"os"
)

func getDayDirectory(day string) string {
	if len(day) == 1 {
		return "./day0" + day + "/"
	}

	return "./day" + day + "/"
}

func getInputFile() string {
	if len(os.Args) == 3 {
		return "puzzle"
	}

	return os.Args[3]
}

func main() {
	if len(os.Args) < 3 {
		panic("Not enough arguments - supply day and problem numbers")
	}

	days := map[string]map[string]util.Solution{
		"1": {
			"1": day1_problem1.Day1Solution1{},
			"2": day1_problem2.Day1Solution2{},
		},
		"2": {
			"1": day2_problem1.Day2Solution1{},
			"2": day2_problem2.Day2Solution2{},
		},
		"3": {
			"1": day3_problem1.Day3Solution1{},
			"2": day3_problem2.Day3Solution2{},
		},
		"4": {
			"1": day4_problem1.Day4Solution1{},
			"2": day4_problem2.Day4Solution2{},
		},
		"5": {
			"1": day5_problem1.Day5Solution1{},
			"2": day5_problem2.Day5Solution2{},
		},
		"6": {
			"1": day6_problem1.Day6Solution1{},
			"2": day6_problem2.Day5Solution2{},
		},
		"7": {
			"1": day7_problem1.Day7Solution1{},
			"2": day7_problem2.Day7Solution2{},
		},
		"8": {
			"1": day8_problem1.Day8Solution1{},
			"2": day8_problem2.Day8Solution2{},
		},
		"9": {
			"1": day9_problem1.Day9Solution1{},
			"2": day9_problem2.Day9Solution2{},
		},
		"10": {
			"1": day10_problem1.Day10Solution1{},
			"2": day10_problem2.Day10Solution2{},
		},
	}

	day := os.Args[1]
	problems, found := days[day]
	if !found {
		panic("Day \"" + day + "\" does not exist.")
	}

	problem := os.Args[2]
	solution, found := problems[problem]
	if !found {
		panic("Problem \"" + problem + "\" does not exist. Must be either 1 or 2.")
	}

	directory := getDayDirectory(day)
	file := getInputFile() + ".input"
	solution.Solve(directory + file)
}
