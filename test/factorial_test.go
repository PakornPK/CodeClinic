package factorial_test

import (
	"testing"
)

func Factorial(nums int) int {
	if nums == 0 {
		return 1
	}
	return nums * Factorial(nums-1)
}

func TestFactorial(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{5, 120},
		{7, 5040},
		{12, 479001600},
	}
	for _, c := range cases {
		got := Factorial(c.in)
		if got != c.want {
			t.Errorf("Factorial(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
