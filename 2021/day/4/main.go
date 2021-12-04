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

	var boards [][][]int64
	var numbersDrawn []int64

	scanner.Scan()
	line := scanner.Text()
	for _, substring := range strings.Split(line, ",") {
		number, err := strconv.ParseInt(substring, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbersDrawn = append(numbersDrawn, number)
	}

	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 0 {
			var board [][]int64

			for i := 0; i < 5; i++ {
				if i > 0 {
					scanner.Scan()
					line = scanner.Text()
				}
				var numbers []int64
				for _, substring := range strings.Fields(line) {
					number, err := strconv.ParseInt(substring, 10, 64)
					if err != nil {
						log.Fatal(err)
					}
					numbers = append(numbers, number)
				}
				board = append(board, numbers)
			}
			boards = append(boards, board)
		}
	}

	var score int64

	for _, number := range numbersDrawn {
		for i, board := range boards {
			for j, line := range board {
				for k, v := range line {
					if number == v {
						boards[i][j][k] = -1
					}
				}
			}
		}

		var bingo bool

		for _, board := range boards {
			var xSums, ySums [5]int64
			var boardSum int64

			for j, line := range board {
				for k, v := range line {
					xSums[j] += v
					ySums[k] += v
				}
			}

			for _, v := range xSums {
				if v < 0 {
					bingo = true
					break
				}
			}

			for _, v := range ySums {
				if v < 0 {
					bingo = true
					break
				}
			}

			if bingo {
				for _, line := range board {
					for _, v := range line {
						if v >= 0 {
							boardSum += v
						}
					}
				}

				score = number * boardSum
				break
			}
		}
		if bingo {
			break
		}
	}
	println(score)

	var keepBoards = boards

	for _, number := range numbersDrawn {
		var tempKeepBoards [][][]int64
		for i, board := range keepBoards {
			for j, line := range board {
				for k, v := range line {
					if number == v {
						keepBoards[i][j][k] = -1
					}
				}
			}
		}

		for _, board := range keepBoards {
			var bingo bool
			var xSums, ySums [5]int64
			var boardSum int64

			for j, line := range board {
				for k, v := range line {
					xSums[j] += v
					ySums[k] += v
				}
			}

			for _, v := range xSums {
				if v < 0 {
					bingo = true
					break
				}
			}

			for _, v := range ySums {
				if v < 0 {
					bingo = true
					break
				}
			}

			if len(keepBoards) == 1 && bingo {
				for _, line := range board {
					for _, v := range line {
						if v >= 0 {
							boardSum += v
						}
					}
				}

				score = number * boardSum

				println(score)
			} else if !bingo {
				tempKeepBoards = append(tempKeepBoards, board)
			}
		}

		keepBoards = tempKeepBoards
	}

}
