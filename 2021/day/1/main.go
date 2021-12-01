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

	var scans []int64

	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		scans = append(scans, value)
	}

	// part 1
	counter := 0

	for k, v := range scans {
		if k > 0 && v > scans[k-1] {
			counter++
		}
	}

	println(counter)

	// part 2
	counter = 0

	for i := 3; i < len(scans); i++ {
		if scans[i] > scans[i-3] { // ignore shared values
			counter++
		}
	}

	println(counter)
}
