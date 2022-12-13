package main

import (
	"reflect"
	"testing"
)

func TestBasicSolution(t *testing.T) {
	actual := maxProfit([]int{7, 1, 5, 3, 6, 4})
	expected := 5
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected : %v, Actual %v", expected, actual)
	}
}
