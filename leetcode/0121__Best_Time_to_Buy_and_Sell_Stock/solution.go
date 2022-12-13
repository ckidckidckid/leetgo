package main

import (
	"math"
)

func max(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}
	hi := math.MinInt
	for _, v := range arr {
		if v > hi {
			hi = v
		}
	}
	return hi
}

func maxProfit(prices []int) int {
	ans, lowest := 0, math.MaxInt
	for _, v := range prices {
		ans = max(ans, v-lowest)
		if v < lowest {
			lowest = v
		}
	}
	return ans
}
