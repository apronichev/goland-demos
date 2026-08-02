package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher", "en")
	want := "Hello, Gopher!"
	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}

func TestSum(t *testing.T) {
	if got := Sum(2, 3); got != 5 {
		t.Errorf("Sum(2, 3) = %d, want 5", got)
	}
}

func TestShout(t *testing.T) {
	got := Shout("hello")
	want := "HELLO!"
	if got != want {
		t.Errorf("Shout(%q) = %q, want %q", "hello", got, want)
	}
}
