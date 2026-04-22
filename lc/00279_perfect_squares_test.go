package algorithms

import (
	"testing"
)

func TestPerfectSquare1(t *testing.T) {
	input := 12

	result := PerfectSquares(input)
	expected := 3

	if result != expected {
		t.Fatalf("failed with input: %v expected: %v result was %v\n", input, expected, result)
	}
}

func TestPerfectSquare2(t *testing.T) {
	input := 13

	result := PerfectSquares(input)
	expected := 2

	if result != expected {
		t.Fatalf("failed with input: %v expected: %v result was %v\n", input, expected, result)
	}
}
