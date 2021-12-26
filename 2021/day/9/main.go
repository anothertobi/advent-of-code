package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	inputFile, err := os.Open("data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var heightmap [][]int

	for scanner.Scan() {
		var heightmapLine []int

		input := scanner.Text()

		for i := 0; i < len(input); i++ {
			number, err := strconv.Atoi(string(input[i]))
			if err != nil {
				log.Fatal(err)
			}

			heightmapLine = append(heightmapLine, number)
		}

		heightmap = append(heightmap, heightmapLine)
	}

	var riskSum int

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			lowest := true
			if (i > 0 && heightmap[i-1][j] <= heightmap[i][j]) ||
				(i < len(heightmap)-1 && heightmap[i+1][j] <= heightmap[i][j]) ||
				(j > 0 && heightmap[i][j-1] <= heightmap[i][j]) ||
				(j < len(heightmap[i])-1 && heightmap[i][j+1] <= heightmap[i][j]) {
				lowest = false
			}

			if lowest {
				riskSum += 1 + heightmap[i][j]
			}
		}
	}

	println(riskSum)

	largestBasins := make([]map[Coordinate]bool, 3)

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			if heightmap[i][j] < 9 {
				basinCoordinates := make(map[Coordinate]bool)

				checkSurrounding(Coordinate{x: j, y: i}, &heightmap, basinCoordinates)

				exists := false

				for _, basin := range largestBasins {
					if reflect.DeepEqual(basinCoordinates, basin) {
						exists = true
						break
					}
					if len(basinCoordinates) == len(basin) {
						println(len(basin))
					}
				}

				if !exists {
					minIndex := minSliceIndex(&largestBasins)
					if len(basinCoordinates) > len(largestBasins[minIndex]) {
						largestBasins[minIndex] = basinCoordinates
					}
				}
			}
		}
	}

	println(len(largestBasins[0]) * len(largestBasins[1]) * len(largestBasins[2]))
}

func checkSurrounding(coordinate Coordinate, heightmap *[][]int, basinCoordinates map[Coordinate]bool) {
	if !basinCoordinates[coordinate] {
		basinCoordinates[coordinate] = true

		if coordinate.y > 0 && (*heightmap)[coordinate.y-1][coordinate.x] < 9 {
			checkSurrounding(Coordinate{coordinate.x, coordinate.y - 1}, heightmap, basinCoordinates)
		}
		if coordinate.y < len((*heightmap))-1 && (*heightmap)[coordinate.y+1][coordinate.x] < 9 {
			checkSurrounding(Coordinate{coordinate.x, coordinate.y + 1}, heightmap, basinCoordinates)
		}
		if coordinate.x > 0 && (*heightmap)[coordinate.y][coordinate.x-1] < 9 {
			checkSurrounding(Coordinate{coordinate.x - 1, coordinate.y}, heightmap, basinCoordinates)
		}
		if coordinate.x < len((*heightmap)[coordinate.y])-1 && (*heightmap)[coordinate.y][coordinate.x+1] < 9 {
			checkSurrounding(Coordinate{coordinate.x + 1, coordinate.y}, heightmap, basinCoordinates)
		}
	}
}

type Coordinate struct {
	x, y int
}

func minSliceIndex(slice *[]map[Coordinate]bool) int {
	min := 0
	minIndex := 0

	for key, value := range *slice {
		if key == 0 {
			min = len(value)
			minIndex = key
		}
		if len(value) <= min {
			min = len(value)
			minIndex = key
		}
	}

	return minIndex
}
