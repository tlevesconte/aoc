package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solveP1() int {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	
	re := regexp.MustCompile(`mul\((-?\d+(\.\d+)?),\s*(-?\d+(\.\d+)?)\)`)
	rr := regexp.MustCompile(`-?\d+(\.\d+)?`)

	var mulSum int
	for scanner.Scan() {
		strs := re.FindAllString(scanner.Text(), -1)
		for _, str := range strs {
			nums := rr.FindAllString(str, -1)

			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])

			mulSum += x * y
		}
	}

	return mulSum
}

func main() {
	fmt.Println("03/P1:", solveP1())
}