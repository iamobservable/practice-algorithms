package algorithms

import (
	"reflect"
	"testing"
)

func TestRotateImageCase1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	RotateImage(matrix)
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	if reflect.DeepEqual(matrix, expected) == false {
		t.Fatalf("failed to pass test case with\nmatrix: %v expected: %v\n", matrix, expected)
	}
}

func TestRotateImageCase2(t *testing.T) {
	matrix := [][]int{
		{4, 8},
		{3, 6},
	}

	RotateImage(matrix)
	expected := [][]int{
		{3, 4},
		{6, 8},
	}

	if reflect.DeepEqual(matrix, expected) == false {
		t.Fatalf("failed to pass test case with\nmatrix: %v expected: %v\n", matrix, expected)
	}
}

func TestRotateImageCase3(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	RotateImage(matrix)
	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	if reflect.DeepEqual(matrix, expected) == false {
		t.Fatalf("failed to pass test case with\nmatrix: %v expected: %v\n", matrix, expected)
	}
}
