package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var left, right []int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		a, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		left = append(left, a)
		right = append(right, b)
	}

	fmt.Println("01/P1:", solveP1(left, right))
	fmt.Println("01/P2:", solveP2(left, right))
}