package main

import (
	"adventofcode/2021/internal/helpers"

	"log"
)

func main() {
	measurements := helpers.GetFileLines("input.txt")

	counter1 := oneDimensionalMeasurements(measurements)
	counter2 := threeDimensionalMeasurements(measurements)

	log.Printf("Amount of increased measurements within one dimensional measurements: %v", counter1)
	log.Printf("Amount of increased measurements within three dimensional measurements: %v", counter2)

}

func oneDimensionalMeasurements(measurements []string) int {
	counter := 0

	for i := range measurements {
		k := len(measurements) - 1 - i

		if k == 0 {
			break
		}

		distance := helpers.ConvertToInt((measurements[k])) - helpers.ConvertToInt((measurements[k-1]))

		if distance > 0 {
			counter = counter + 1
		}
	}

	return counter
}

func threeDimensionalMeasurements(measurements []string) int {
	counter := 0

	for i := range measurements {
		k := len(measurements) - 1 - i
		m := k - 1

		if m-2 <= 0 || k-2 <= 0 {
			break
		}

		latestThree := helpers.ConvertToInt((measurements[k])) + helpers.ConvertToInt((measurements[k-1])) + helpers.ConvertToInt((measurements[k-2]))
		previousThree := helpers.ConvertToInt((measurements[m])) + helpers.ConvertToInt((measurements[m-1])) + helpers.ConvertToInt((measurements[m-2]))

		distance := latestThree - previousThree

		if distance > 0 {
			counter = counter + 1
		}
	}

	return counter
}
