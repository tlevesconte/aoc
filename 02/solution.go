package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveP1() int {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var safeCount int
	for scanner.Scan() {
		line := toInt(strings.Fields(scanner.Text()))

		if safe(line) {
			safeCount++
		}
	}

	return safeCount
}

func solveP2() int {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var safeCount int
	for scanner.Scan() {
		line := toInt(strings.Fields(scanner.Text()))

		if safe(line) || dampen(line) {
			safeCount++
		}
	}

	return safeCount
}

func toInt(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		x, _ := strconv.Atoi(str)
		ints[i] = x
	}

	return ints
}

func safe(input []int) bool {
	incr := input[0] < input[1]

	for i := 1; i < len(input); i++ {
		if incr != (input[i-1] < input[i]) {
			return false
		}

		if input[i-1] == input[i] {
			return false
		}

		if abs(input[i-1], input[i]) > 3 {
			return false
		}
	}

	return true
}

func dampen(input []int) bool {
	for i := 0; i < len(input); i++ {
		ints := append([]int{}, input[:i]...)
		ints = append(ints, input[i+1:]...)

		if safe(ints) {
			return true
		}
	}

	return false
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

func main() {
	fmt.Println("02/P1:", solveP1())
	fmt.Println("02/P2:", solveP2())
}