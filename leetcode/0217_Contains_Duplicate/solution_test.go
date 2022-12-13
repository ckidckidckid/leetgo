package main

import (
	"reflect"
	"testing"
)

func TestBasicSolution1(t *testing.T) {
	actual := containsDuplicate([]int{1, 2, 3, 1})
	expected := true
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected : %v, Actual %v", expected, actual)
	}
}

func TestBasicSolution2(t *testing.T) {
	actual := containsDuplicate([]int{1, 2, 3, 4})
	expected := false
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected : %v, Actual %v", expected, actual)
	}
}
