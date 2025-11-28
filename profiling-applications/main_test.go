package main

import "testing"

func TestFib2(t *testing.T) {
	if FibRecursive(40) != 102334155 {
		t.Error("Incorrect!")
	}
}
