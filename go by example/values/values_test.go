package main

import "testing"

func TestConcatenateStrings(t *testing.T) {
	got := ConcatenateStrings("Hello, ", "World!")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAddIntegers(t *testing.T) {
	got := AddIntegers(1, 2)
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDivideFloats(t *testing.T) {
	got := DivideFloats(7, 3)
	want := 2.33

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestAnd(t *testing.T) {
	got := And(true, true)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestOr(t *testing.T) {
	got := Or(true, false)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
func TestNot(t *testing.T) {
	got := Not(true)
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
