package main

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for idx, val := range nums {
		otherIdx, isPresent := table[target-nums[idx]]
		if isPresent {
			return []int{otherIdx, idx}
		} else {
			table[val] = idx
		}
	}
	return []int{}
}
