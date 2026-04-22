package algorithms

import (
	"testing"
)

func TestReverseStringRecursively(t *testing.T) {
	input := "ken thompson"
	expected := "nospmoht nek"
	output := RecursiveReverse(input)

	if output != expected {
		t.Fatalf("failed to sort string recursively, instead returned \"%v\"\n", output)
	}
}

func TestReverseStringProcedurally(t *testing.T) {
	input := "ken thompson"
	expected := "nospmoht nek"
	output := ProceduralReverse(input)

	if output != expected {
		t.Fatalf("failed to sort string procedurally, instead returned \"%v\"\n", output)
	}
}
