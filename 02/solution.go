package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var mat [][]int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		var vec []int
		for _, str := range line {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}

			vec = append(vec, num)
		}
		mat = append(mat, vec)
	}

	fmt.Println("02/P1:", solveP1(mat))
}

func solveP1(input [][]int) int {
	var safeCount int 
	for _, vec := range input {
		isIncreasing := vec[0] < vec[1]
		isSafe := true

		for i := 1; i < len(vec); i++ {
			if isIncreasing != (vec[i-1] < vec[i]) {
				isSafe = false
				break
			}
			if vec[i-1] == vec[i] {
				isSafe = false
				break
			}
			if absDiff(vec[i-1], vec[i]) > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}