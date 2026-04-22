package algorithms

import (
	"testing"
)

func TestGasStationCase1(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	result := CanCompleteCircuit(gas, cost)
	expected := 3

	if result != expected {
		t.Fatalf("failed to pass test case with\ngas: %v and cost: %v expected: %v, result: %v\n", gas, cost, expected, result)
	}
}

func TestGasStationCase2(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	result := CanCompleteCircuit(gas, cost)
	expected := -1

	if result != expected {
		t.Fatalf("failed to pass test case with\ngas: %v and cost: %v expected: %v, result: %v\n", gas, cost, expected, result)
	}
}

func TestGasStationCase3(t *testing.T) {
	gas := []int{5, 1, 2, 3, 4}
	cost := []int{4, 4, 1, 5, 1}

	result := CanCompleteCircuit(gas, cost)
	expected := 4

	if result != expected {
		t.Fatalf("failed to pass test case with\ngas: %v and cost: %v expected: %v, result: %v\n", gas, cost, expected, result)
	}
}
