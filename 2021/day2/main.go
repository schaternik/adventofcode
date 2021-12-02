package main

import (
	"adventofcode/2021/internal/helpers"

	"log"
	"strings"
)

const UP = "up"
const DOWN = "down"
const FORWARD = "forward"

func main() {
	positions := helpers.GetFileLines("input.txt")

	course1 := course1(positions)
	course2 := course2(positions)

	log.Printf("Course 1: %v", course1)
	log.Printf("Course 2: %v", course2)

}

func course1(positions []string) int {
	horizontalPosition := 0
	depth := 0

	for _, p := range positions {
		a := strings.Split(p, " ")
		action, points := a[0], a[1]

		switch action {
		case UP:
			depth -= helpers.ConvertToInt(points)
		case DOWN:
			depth += helpers.ConvertToInt(points)
		case FORWARD:
			horizontalPosition += helpers.ConvertToInt(points)
		}
	}

	return horizontalPosition * depth
}

func course2(positions []string) int {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, p := range positions {
		a := strings.Split(p, " ")
		action, points := a[0], a[1]

		switch action {
		case UP:
			aim -= helpers.ConvertToInt(points)
		case DOWN:
			aim += helpers.ConvertToInt(points)
		case FORWARD:
			horizontalPosition += helpers.ConvertToInt(points)
			depth += (aim * helpers.ConvertToInt(points))
		}
	}

	return horizontalPosition * depth
}
