package algorithms

import (
	"testing"
)

func TestCanPlaceFlowersCase1(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1

	result := CanPlaceFlowers(flowerbed, n)
	expected := true

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase2(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2

	result := CanPlaceFlowers(flowerbed, n)
	expected := false

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase3(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2

	result := CanPlaceFlowers(flowerbed, n)
	expected := false

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase4(t *testing.T) {
	flowerbed := []int{1, 0, 1, 0, 1, 0, 1}
	n := 1

	result := CanPlaceFlowers(flowerbed, n)
	expected := false

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase5(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1, 0, 0}
	n := 2

	result := CanPlaceFlowers(flowerbed, n)
	expected := true

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase6(t *testing.T) {
	flowerbed := []int{0}
	n := 1

	result := CanPlaceFlowers(flowerbed, n)
	expected := true

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase7(t *testing.T) {
	flowerbed := []int{1}
	n := 1

	result := CanPlaceFlowers(flowerbed, n)
	expected := false

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase8(t *testing.T) {
	flowerbed := []int{0, 1}
	n := 1

	result := CanPlaceFlowers(flowerbed, n)
	expected := false

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}

func TestCanPlaceFlowersCase9(t *testing.T) {
	flowerbed := []int{0, 1, 0}
	n := 1

	result := CanPlaceFlowers(flowerbed, n)
	expected := false

	if result != expected {
		t.Fatalf("failed to pass test case with\nflowerbed: %v and n: %v expected: %v, result: %v\n", flowerbed, n, expected, result)
	}
}
