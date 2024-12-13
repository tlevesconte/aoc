package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solveP1(input string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	ms := re.FindAllStringSubmatch(input, -1)

	mulSum := 0
	for _, m := range ms {
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])

		mulSum += x * y
	}

	return mulSum
}

func solveP2(input string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	ms := re.FindAllStringSubmatch(input, -1)

	mulSum := 0
	enabled := true
	for _, m := range ms {
		switch m[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				x, _ := strconv.Atoi(m[1])
				y, _ := strconv.Atoi(m[2])

				mulSum += x * y
			}
		}
	}

	return mulSum
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var text string
	for scanner.Scan() {
		text += scanner.Text()
	}

	fmt.Println("03/P1:", solveP1(text))
	fmt.Println("03/P2:", solveP2(text))
}
