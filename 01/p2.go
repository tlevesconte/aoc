package main

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
