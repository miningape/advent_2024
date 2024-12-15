package day14

import (
	"advent2024/util"
	"strings"
)

type Matcher struct {
	*util.Matcher
}

type Robot struct {
	Position util.Vector
	Velocity util.Vector
}

func (matcher *Matcher) matchRobot() Robot {
	robot := Robot{}

	matcher.Match("p=")
	robot.Position.X = matcher.Match_int()
	matcher.Match(",")
	robot.Position.Y = matcher.Match_int()

	matcher.Match(" v=")
	robot.Velocity.X = matcher.Match_int()
	matcher.Match(",")
	robot.Velocity.Y = matcher.Match_int()

	return robot
}

func (matcher *Matcher) matchRobots() []Robot {
	robots := make([]Robot, 0, len(strings.Split(matcher.Source, "\n")))

	for !matcher.IsAtEnd() {
		robot := matcher.matchRobot()
		robots = append(robots, robot)

		if !matcher.IsAtEnd() {
			matcher.Match("\n")
		}
	}

	return robots
}

func ParseInput(file string) []Robot {
	matcher := Matcher{&util.Matcher{Source: file, Index: 0}}
	return matcher.matchRobots()
}