package main

import (
	"testing"
)

func TestRoman(t *testing.T) {
	answer := Roman("MMMDCCXXIV")
	if answer != 3726 {
		t.Fatalf(`Answer is 3726, you got answer=%v`, answer)
	} //if test fails
}

func TestDupRemover(t *testing.T) {

	var array = removeDuplicates([]int{1, 2, 3, 3, 3, 5, 7, 8, 9, 9})
	var correctArray = []int{1, 2, 3, 5, 7, 8, 9}

	for i := 0; i < len(array); i++ {
		if array[i] != correctArray[i] {
			t.Fatalf(`Expected correctArray=%v, you got array=%v`, correctArray, array)
		}
	}
}

func TestDupConstraintDecreasingOrder(t *testing.T) {
	var array = removeDuplicates([]int{1, 2, 3, 3, 5, 3, 7, 8, 9, 9})
	if len(array) != 0 {
		t.Fatalf(`Expected empty array, you got array=%v`, array)
	}
}

func TestDupConstraintLargeArrayNumber(t *testing.T) {
	var array = removeDuplicates([]int{1, 2, 3, 3, 3, 7, 8, 9, 999})
	if len(array) != 0 {
		t.Fatalf(`Expected empty array, you got array=%v`, array)
	}
}
