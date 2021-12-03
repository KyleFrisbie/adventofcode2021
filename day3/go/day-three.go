package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := readInput()
	inputs := scanFile(inputFile)
	sums := scanMeasurements(inputs)
	getRates(sums, len(inputs))
}

func readInput() string {
	var inputFile string
	flag.StringVar(&inputFile, "i", "../input-test.txt", "Input values file.")
	flag.Parse()
	return inputFile
}

func scanFile(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var measurements [][]int
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		var b []int
		for _, i := range s {
			b = append(b, stringToInt(i))
		}
		measurements = append(measurements, b)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return measurements
}

func scanMeasurements(measurements [][]int) []int {
	x := make([]int, len(measurements[0]))
	for i := 0; i < len(x); i++ {
		x[i] = 0
	}
	for _, m := range measurements {
		for i := 0; i < len(m); i++ {
			x[i] += m[i]
		}
	}
	return x
}

func getRates(sums []int, totalMeasurements int) (int, int) {
	gammaRate := make([]int, len(sums))
	epsilonRate := make([]int, len(sums))

	for i := 0; i < len(sums); i++ {
		avg := float64(sums[i]) / float64(totalMeasurements)
		gammaRate[i] = int(math.Floor(avg + 0.5))

		epsilonRate[i] = 0
		if gammaRate[i] == 0 {
			epsilonRate[i] = 1
		}
	}
	gr := binaryAsIntArrayToInt(gammaRate)
	er := binaryAsIntArrayToInt(epsilonRate)

	fmt.Println(gr)
	fmt.Println(er)
	return gr, er
}

func binaryAsIntArrayToInt(x []int) int {
	var vals string
	for _, i := range x {
		vals += strconv.Itoa(i)
	}
	num, err := strconv.Atoi(vals)
	if err != nil {
		fmt.Printf("Unable to convert vals to num: %s", err)
	}
	s, _ := strconv.ParseUint(vals, 10, 32)
	fmt.Printf("num: %d\n", s)
	var b bytes.Buffer
	b.w
	value := fmt.Sprint(b.Bytes())
	fmt.Println(value)
	v, _ := strconv.Atoi(value)
	return v
}

func stringToInt(text string) int {
	text = strings.Replace(text, " ", "", -1)
	i, err := strconv.Atoi(text)

	if err != nil {
		fmt.Printf("Error converting text to int.")
	}
	return i
}
