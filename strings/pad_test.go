package strings

import "testing"

func TestPadLeft(t *testing.T) {
	s := PadLeft("1", "abc", 3)
	t.Logf("s: %s\n", s)
}

func TestPadRight(t *testing.T) {
	s := PadRight("1", "abc", 3)
	t.Logf("s: %s\n", s)
}
