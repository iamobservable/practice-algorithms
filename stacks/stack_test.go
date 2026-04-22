package algorithms

import (
	"testing"
)

func TestStackCannotPushWhenZeroSize(t *testing.T) {
	stack := NewStack(0)

	if stack.Push("11") {
		t.Fatalf("failed to push into a zero sized Stack\n")
	}
}

func TestStackCanPushWhenEmpty(t *testing.T) {
	stack := NewStack(2)

	if stack.Push("first") == false {
		t.Fatalf("failed to push when empty")
	}
}

func TestStackCannotPushWhenFull(t *testing.T) {
	stack := NewStack(2)
	stack.Push("first")
	stack.Push("second")

	if stack.Push("third") == true {
		t.Fatalf("failed to limit push when full")
	}
}

func TestStackDisplaysCorrectPeekValues(t *testing.T) {
	stack := NewStack(3)
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	seen := stack.Peek()

	if seen != "third" {
		t.Fatalf("failed display current value \"%v\" with peek, returned %v instead\n", "first", seen)
	}

	stack.Pop()

	if seen != "second" {
		t.Fatalf("failed display current value \"%v\" with peek, returned %v instead\n", "second", seen)
	}

	stack.Pop()

	if seen != "third" {
		t.Fatalf("failed display current value \"%v\" with peek, returned %v instead\n", "third", seen)
	}
}

func TestStackIsEmptyAfterPushAndPop(t *testing.T) {
	stack := NewStack(3)
	stack.Push("first")
	stack.Pop()

	if stack.IsEmpty() == false {
		t.Fatalf("failed to return correct value of IsEmpty")
	}
}

func TestStackIsNotEmptyAfterPushPushAndPop(t *testing.T) {
	stack := NewStack(3)
	stack.Push("first")
	stack.Push("second")
	stack.Pop()

	if stack.IsEmpty() {
		t.Fatalf("failed to return correct value of IsEmpty")
	}
}
