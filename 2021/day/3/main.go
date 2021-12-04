package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var counters []int

	for i := 0; i < 12; i++ {
		counter := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '1' {
				counter++
			} else {
				counter--
			}
		}
		counters = append(counters, counter)
	}

	var gamma, epsilon string

	for i := 0; i < len(counters); i++ {
		if counters[i] >= 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaRate, err := strconv.ParseUint(gamma, 2, 12)
	if err != nil {
		log.Fatal(err)
	}

	epsilonRate, err := strconv.ParseUint(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	println(gammaRate * epsilonRate)

	oxygen, err := getOxygenFromLines(lines)
	if err != nil {
		log.Fatal(err)
	}

	co2, err := getCO2FromLines(lines)
	if err != nil {
		log.Fatal(err)
	}

	println(oxygen * co2)
}

func getOxygenFromLines(lines []string) (uint64, error) {
	var keepNumbers = lines

	for i := 0; len(keepNumbers) != 1; i++ {
		var tempKeepNumbers []string
		var counter int

		for j := 0; j < len(keepNumbers); j++ {
			if keepNumbers[j][i] == '1' {
				counter++
			} else {
				counter--
			}
		}

		for _, v := range keepNumbers {
			if counter >= 0 && v[i] == '1' {
				tempKeepNumbers = append(tempKeepNumbers, v)
			} else if counter < 0 && v[i] == '0' {
				tempKeepNumbers = append(tempKeepNumbers, v)
			}
		}

		keepNumbers = tempKeepNumbers
	}

	oxygen, err := strconv.ParseUint(keepNumbers[0], 2, 64)
	if err != nil {
		return 0, err
	}

	return oxygen, nil
}

func getCO2FromLines(lines []string) (uint64, error) {
	var keepNumbers = lines

	for i := 0; len(keepNumbers) != 1; i++ {
		var tempKeepNumbers []string
		var counter int

		for j := 0; j < len(keepNumbers); j++ {
			if keepNumbers[j][i] == '1' {
				counter++
			} else {
				counter--
			}
		}

		for _, v := range keepNumbers {
			if counter >= 0 && v[i] == '0' {
				tempKeepNumbers = append(tempKeepNumbers, v)
			} else if counter < 0 && v[i] == '1' {
				tempKeepNumbers = append(tempKeepNumbers, v)
			}
		}

		keepNumbers = tempKeepNumbers
	}

	co2, err := strconv.ParseUint(keepNumbers[0], 2, 64)
	if err != nil {
		return 0, err
	}

	return co2, nil
}
