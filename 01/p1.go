package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(solve(file))
}

func solve(input *os.File) int {
	scanner := bufio.NewScanner(input)

	var left, right []int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])

		left = append(left, a)
		right = append(right, b)
	}

	slices.Sort(left)
	slices.Sort(right)

	var diff int
	for i := range left {
		diff += abs(left[i], right[i])
	}

	return diff
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}