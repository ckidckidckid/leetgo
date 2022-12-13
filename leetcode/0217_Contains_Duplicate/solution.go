package main

func containsDuplicate(nums []int) bool {
	table := make(map[int]struct{})
	for _, val := range nums {
		if _, exists := table[val]; exists {
			return true
		} else {
			table[val] = struct{}{}
		}
	}
	return false
}
