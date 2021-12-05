package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type Line struct {
	StartCoordinate, EndCoordinate Coordinate
}

func main() {
	inputFile, err := os.Open("data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var lines []Line
	var counter int
	coordinates := map[Coordinate]int{}

	for scanner.Scan() {
		input := scanner.Text()

		var line Line

		substrings := strings.Split(input, " -> ")

		line.StartCoordinate, err = GetCoordinateFromString(substrings[0])
		if err != nil {
			log.Fatal(err)
		}
		line.EndCoordinate, err = GetCoordinateFromString(substrings[1])
		if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, line)
	}

	for _, line := range lines {
		var delta int

		deltaX := line.EndCoordinate.X - line.StartCoordinate.X
		deltaY := line.EndCoordinate.Y - line.StartCoordinate.Y

		if deltaX != 0 {
			delta = deltaX
		} else {
			delta = deltaY
		}

		if delta < 0 {
			delta = -delta
		}

		for i := 0; i <= delta; i++ {
			coordinate := line.StartCoordinate

			if deltaX > 0 {
				coordinate.X += i
			} else if deltaX < 0 {
				coordinate.X -= i
			}
			if deltaY > 0 {
				coordinate.Y += i
			} else if deltaY < 0 {
				coordinate.Y -= i
			}

			coordinates[coordinate]++
			if coordinates[coordinate] == 2 {
				counter++
			}
		}
	}

	println(counter)
}

func GetCoordinateFromString(input string) (Coordinate, error) {
	var coordinate Coordinate

	substrings := strings.Split(input, ",")

	number, err := strconv.Atoi(substrings[0])
	if err != nil {
		return coordinate, err
	}
	coordinate.X = int(number)

	number, err = strconv.Atoi(substrings[1])
	if err != nil {
		return coordinate, err
	}
	coordinate.Y = number

	return coordinate, err
}
