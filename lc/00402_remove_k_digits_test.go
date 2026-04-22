package algorithms

import (
	"testing"
)

func TestRemoveKDigitsCase1(t *testing.T) {
	num := "1432219"
	k := 3

	result := RemoveKDigits(num, k)
	expected := "1219"

	if result != expected {
		t.Fatalf("failed with num: %v k: %v expected: %v result was %v\n", num, k, expected, result)
	}
}

func TestRemoveKDigitsCase2(t *testing.T) {
	num := "10200"
	k := 1

	result := RemoveKDigits(num, k)
	expected := "200"

	if result != expected {
		t.Fatalf("failed with num: %v k: %v expected: %v result was %v\n", num, k, expected, result)
	}
}

func TestRemoveKDigitsCase3(t *testing.T) {
	num := "10"
	k := 2

	result := RemoveKDigits(num, k)
	expected := "0"

	if result != expected {
		t.Fatalf("failed with num: %v k: %v expected: %v result was %v\n", num, k, expected, result)
	}
}

func TestRemoveKDigitsCase4(t *testing.T) {
	num := "10"
	k := 1

	result := RemoveKDigits(num, k)
	expected := "0"

	if result != expected {
		t.Fatalf("failed with num: %v k: %v expected: %v result was %v\n", num, k, expected, result)
	}
}

func TestRemoveKDigitsCase5(t *testing.T) {
	num := "1111111"
	k := 3

	result := RemoveKDigits(num, k)
	expected := "1111"

	if result != expected {
		t.Fatalf("failed with num: %v k: %v expected: %v result was %v\n", num, k, expected, result)
	}
}

func TestRemoveKDigitsCase6(t *testing.T) {
	num := "112"
	k := 1

	result := RemoveKDigits(num, k)
	expected := "11"

	if result != expected {
		t.Fatalf("failed with num: %v k: %v expected: %v result was %v\n", num, k, expected, result)
	}
}

func TestRemoveKDigitsCase7(t *testing.T) {
	num := "12345"
	k := 2

	result := RemoveKDigits(num, k)
	expected := "145"

	if result != expected {
		t.Fatalf("failed with num: %v k: %v expected: %v result was %v\n", num, k, expected, result)
	}
}
