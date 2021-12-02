package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ConvertToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Error when converting to integer: %s", s)
	}

	return n
}

func GetFileLines(fileName string) []string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)

	return strings.Split(input, "\n")
}
