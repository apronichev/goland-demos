package greeting

import "testing"

func TestGreeterGreet(t *testing.T) {
	got := Default().Greet("Gopher")
	want := "Hello, Gopher!"
	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}

func TestShout(t *testing.T) {
	if got := Shout("hi"); got != "HI" {
		t.Errorf("Shout() = %q, want %q", got, "HI")
	}
}
