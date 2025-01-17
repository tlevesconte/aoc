package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solveP1(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	var diff int
	for i := range left {
		diff += abs(left[i], right[i])
	}

	return diff
}

func solveP2(left, right []int) int {
	simMap := make(map[int]int)
	for _, val := range right {
		simMap[val]++
	}

	var simScore int
	for _, val := range left {
		simScore += val * simMap[val]
	}

	return simScore
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var left, right []int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])

		left = append(left, a)
		right = append(right, b)
	}

	fmt.Println("01/P1:", solveP1(left, right))
	fmt.Println("01/P2:", solveP2(left, right))
}
