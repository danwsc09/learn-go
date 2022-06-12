package main

import (
	"testing"
)

// assume arr1 and arr2 are sorted
func isSameSortedArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestCalculateHand(t *testing.T) {
	hand := []string{"H4", "S4", "CA"}
	res := CalculateHand(hand)
	expected := []int{9, 19}
	if !isSameSortedArray(res, expected) {
		t.Fatalf("Failed")
	}
}

func TestCalculateHand2(t *testing.T) {
	hand := []string{"C7", "H5"}
	res := CalculateHand(hand)
	expected := []int{12}
	if !isSameSortedArray(res, expected) {
		t.Fatalf("Failed")
	}
}

func TestCalculateHand3(t *testing.T) {
	hand := []string{"C7", "H5", "SJ"}
	res := CalculateHand(hand)
	expected := []int{}
	if !isSameSortedArray(res, expected) {
		t.Fatalf("Failed")
	}
}

func TestCalculateHand4(t *testing.T) {
	hand := []string{"CA", "H10"}
	res := CalculateHand(hand)
	expected := []int{11, 21}
	if !isSameSortedArray(res, expected) {
		t.Fatalf("Failed")
	}
}
