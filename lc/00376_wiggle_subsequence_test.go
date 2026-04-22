package algorithms

import (
	"testing"
)

func TestWiggleSubsequenceCase1(t *testing.T) {
	nums := []int{1, 7, 4, 9, 2, 5}

	result := WiggleSubsequence(nums)
	expected := 6

	if result != expected {
		t.Fatalf("failed with nums: %v expected: %v result was %v\n", nums, expected, result)
	}
}

func TestWiggleSubsequenceCase2(t *testing.T) {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}

	result := WiggleSubsequence(nums)
	expected := 7

	if result != expected {
		t.Fatalf("failed with nums: %v expected: %v result was %v\n", nums, expected, result)
	}
}

func TestWiggleSubsequenceCase3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := WiggleSubsequence(nums)
	expected := 2

	if result != expected {
		t.Fatalf("failed with nums: %v expected: %v result was %v\n", nums, expected, result)
	}
}

func TestWiggleSubsequenceCase4(t *testing.T) {
	nums := []int{84}

	result := WiggleSubsequence(nums)
	expected := 1

	if result != expected {
		t.Fatalf("failed with nums: %v expected: %v result was %v\n", nums, expected, result)
	}
}

func TestWiggleSubsequenceCase5(t *testing.T) {
	nums := []int{0, 0}

	result := WiggleSubsequence(nums)
	expected := 1

	if result != expected {
		t.Fatalf("failed with nums: %v expected: %v result was %v\n", nums, expected, result)
	}
}

func TestWiggleSubsequenceCase6(t *testing.T) {
	nums := []int{0, 0, 0}

	result := WiggleSubsequence(nums)
	expected := 1

	if result != expected {
		t.Fatalf("failed with nums: %v expected: %v result was %v\n", nums, expected, result)
	}
}
