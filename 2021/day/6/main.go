package main

import (
	"bufio"
	"log"
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

	var lanternfishes []int
	lanternfishStateGroups := make([]int, 9)

	scanner.Scan()
	input := scanner.Text()
	lanternfishStates := strings.Split(input, ",")

	for _, lanternfishState := range lanternfishStates {
		number, err := strconv.Atoi(lanternfishState)
		if err != nil {
			log.Fatal(err)
		}

		lanternfishStateGroups[number]++

		lanternfishes = append(lanternfishes, number)
	}

	for i := 0; i < 80; i++ {
		for key, lanternfish := range lanternfishes {
			if lanternfish == 0 {
				lanternfishes[key] = 6
				lanternfishes = append(lanternfishes, 8)
			} else {
				lanternfishes[key]--
			}
		}
	}

	println(len(lanternfishes))

	var sum int

	for i := 0; i < 256; i++ {
		newFish := lanternfishStateGroups[0]
		for j := 0; j < len(lanternfishStateGroups)-1; j++ {
			lanternfishStateGroups[j] = lanternfishStateGroups[j+1]
		}
		lanternfishStateGroups[6] += newFish
		lanternfishStateGroups[8] = newFish
	}

	for _, v := range lanternfishStateGroups {
		sum += v
	}

	println(sum)

}
