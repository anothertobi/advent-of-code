package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const UP = "up"
const DOWN = "down"

type Record struct {
	action string
	value  int64
}

func main() {
	inputFile, err := os.Open("data/input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var records []Record

	for scanner.Scan() {
		var record Record

		fields := strings.Fields(scanner.Text())

		record.action = fields[0]
		record.value, err = strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
	}

	var horizontal, depth, aim int64

	for _, record := range records {
		if record.action == UP {
			depth -= record.value
		} else if record.action == DOWN {
			depth += record.value
		} else {
			horizontal += record.value
		}
	}

	println(depth * horizontal)

	horizontal, depth = 0, 0

	for _, record := range records {
		if record.action == UP {
			aim -= record.value
		} else if record.action == DOWN {
			aim += record.value
		} else {
			horizontal += record.value
			depth += aim * record.value
		}
	}

	println(depth * horizontal)

}
