package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var crabPositions []int

	scanner.Scan()
	input := scanner.Text()
	crabPositionsInput := strings.Split(input, ",")

	for _, crabPosition := range crabPositionsInput {
		number, err := strconv.Atoi(crabPosition)
		if err != nil {
			log.Fatal(err)
		}

		crabPositions = append(crabPositions, number)
	}

	minCrabPosition := crabPositions[0]
	maxCrabPosition := crabPositions[0]

	for _, v := range crabPositions {
		if v < minCrabPosition {
			minCrabPosition = v
		}
		if v > maxCrabPosition {
			maxCrabPosition = v
		}
	}

	fuelRequired := make(map[int]int)
	min := math.MaxInt

	for i := minCrabPosition; i <= maxCrabPosition; i++ {
		for _, v := range crabPositions {
			fuelRequired[i] += int(math.Abs(float64(i) - float64(v)))
		}
	}

	for _, v := range fuelRequired {
		if v < min {
			min = v
		}
	}

	println(min)

	fuelRequired = make(map[int]int)
	min = math.MaxInt

	for i := minCrabPosition; i <= maxCrabPosition; i++ {
		for _, v := range crabPositions {
			delta := int(math.Abs(float64(i) - float64(v)))
			for j := 1; j <= delta; j++ {
				fuelRequired[i] += j
			}
		}
	}

	for _, v := range fuelRequired {
		if v < min {
			min = v
		}
	}

	println(min)
}
