package fuzz_demo

import "testing"

// todo run TestReverse to test the Reverse function. It is successfully finished.
func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello world", "olleh dlrow"},
		{"foo bar", "oof rab"},
		{"Go is fun", "oG si nuf"},
	}

	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// todo run FuzzReverse to test the Reverse function. If there is a bug in the Reverse function,
// todo the fuzz test will eventually find an input that causes the function to fail. For example,
// todo if the Reverse function does not handle certain Unicode characters or multi-byte sequences correctly,
// todo the fuzz test might uncover this issue.
func FuzzReverse(f *testing.F) {
	testcases := []string{"hello world", "foo bar", "Go is fun"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, s string) {
		// Check if Reverse can handle any input without panicking or returning incorrect results
		reversed := Reverse(s)
		doubleReversed := Reverse(reversed)
		if s != doubleReversed {
			t.Errorf("Before: %q, after: %q", s, doubleReversed)
		}
	})
}
