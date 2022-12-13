package main

import (
	"reflect"
	"testing"
)

func TestBasicSolution(t *testing.T) {
	actual := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected : %v, Actual %v", expected, actual)
	}
}
