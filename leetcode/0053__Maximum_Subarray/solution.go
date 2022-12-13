package main

import "math"

func maxSubArray(nums []int) int {
	curr, best := 0, math.MinInt
	for _, n := range nums {
		curr = max(n, curr+n)
		best = max(best, curr)
	}
	return best
}

func max(args ...int) int {
	hi := math.MinInt
	for _, v := range args {
		if v > hi {
			hi = v
		}
	}
	return hi
}
