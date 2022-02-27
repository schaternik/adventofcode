package main

import (
	"adventofcode/2021/internal/helpers"
	"log"
	"strings"

	"github.com/thoas/go-funk"
)

func main() {
	values := helpers.GetFileLines("input.txt")
	valuesMatrix := [][]string{}

	for i := range values {
		valuesMatrix = append(valuesMatrix, strings.Split(values[i], ""))
	}

	gammaRate := []string{}
	epsilonRate := []string{}

	for i := range valuesMatrix {
		verticalSlice := []string{}

		for j := range valuesMatrix[i] {
			verticalSlice = append(verticalSlice, valuesMatrix[i][j])
		}

		foundOnes := funk.FilterString(verticalSlice, func(x string) bool {
			return x == "1"
		})

		foundZeros := funk.FilterString(verticalSlice, func(x string) bool {
			return x == "0"
		})

		if (len(foundOnes)) > len(foundZeros) {
			gammaRate = append(gammaRate, "1")
			epsilonRate = append(epsilonRate, "0")
		} else {
			gammaRate = append(gammaRate, "0")
			epsilonRate = append(epsilonRate, "1")
		}
	}

	gammaRateDecimal := helpers.BinaryConvertToInt(strings.Join(gammaRate, ""))
	epsilonRateDecimal := helpers.BinaryConvertToInt(strings.Join(epsilonRate, ""))

	log.Printf("Power consumption: %v", gammaRateDecimal*epsilonRateDecimal)
	log.Printf("Equal to the right result: %v", gammaRateDecimal*epsilonRateDecimal == 198)
}
