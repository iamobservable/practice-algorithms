package algorithms

import (
	"testing"
)

func TestCandyCase1(t *testing.T) {
	ratings := []int{1, 0, 2}

	result := Candy(ratings)
	expected := 5

	if result != expected {
		t.Fatalf("failed with ratings: %v expected: %v result was %v\n", ratings, expected, result)
	}
}

func TestCandyCase2(t *testing.T) {
	ratings := []int{1, 2, 2}

	result := Candy(ratings)
	expected := 4

	if result != expected {
		t.Fatalf("failed with ratings: %v expected: %v result was %v\n", ratings, expected, result)
	}
}

func TestCandyCase3(t *testing.T) {
	ratings := []int{1, 3, 2, 2, 1}

	result := Candy(ratings)
	expected := 7

	if result != expected {
		t.Fatalf("failed with ratings: %v expected: %v result was %v\n", ratings, expected, result)
	}
}

func TestCandyCase4(t *testing.T) {
	ratings := []int{1, 2, 87, 87, 87, 2, 1}

	result := Candy(ratings)
	expected := 13

	if result != expected {
		t.Fatalf("failed with ratings: %v expected: %v result was %v\n", ratings, expected, result)
	}
}

func TestCandyCase5(t *testing.T) {
	ratings := []int{1, 3, 4, 5, 2}

	result := Candy(ratings)
	expected := 11

	if result != expected {
		t.Fatalf("failed with ratings: %v expected: %v result was %v\n", ratings, expected, result)
	}
}
