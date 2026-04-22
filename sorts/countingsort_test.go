package algorithms

import (
	"reflect"
	"testing"
)

func TestCountingsortAscending(t *testing.T) {
	input := []int{1, 5, 5, 2, 8, 7, 9, 1, 4, 2, 6}
	expected := []int{1, 1, 2, 2, 4, 5, 5, 6, 7, 8, 9}

	sorted := Countingsort(input)

	if reflect.DeepEqual(sorted, expected) == false {
		t.Fatalf("failed to countingsort in an ascending order: %v\n", sorted)
	}
}
