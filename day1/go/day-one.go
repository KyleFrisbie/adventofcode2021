package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, avgRange := readInput()
	increaseCount := scanFile(inputFile)
	countDepthIncrease(increaseCount, avgRange)
}

func readInput() (string, int) {
	var inputFile string
	var averageRange int
	flag.StringVar(&inputFile, "i", "../input-test.txt", "Input values file.")
	flag.IntVar(&averageRange, "a", 1, "Range of measurements to average for calculating depth increase.")
	flag.Parse()
	return inputFile, averageRange
}

func scanFile(inputFile string) []int {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var depths []int
	for scanner.Scan() {
		depths = append(depths, stringToInt(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return depths
}

func countDepthIncrease(measurements []int, avgRange int) {
	lastVal := -1
	currVal := -1
	increases := -1

	for i := (avgRange - 1); i < len(measurements); i++ {
		lastVal = currVal
		currVal = 0
		for j := 0; j < avgRange; j++ {
			currVal += measurements[i-j]
		}

		if currVal > lastVal {
			increases++
		}
	}

	fmt.Printf("Total Increases: %d", increases)
}

func stringToInt(text string) int {
	text = strings.Replace(text, " ", "", -1)
	i, err := strconv.Atoi(text)

	if err != nil {
		fmt.Printf("Error converting text to int.")
	}
	return i
}
