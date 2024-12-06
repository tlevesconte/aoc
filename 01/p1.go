package main

import "slices"

func solveP1(left, right []int) int {
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
