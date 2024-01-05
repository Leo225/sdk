package random

import "testing"

func TestNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("Number: %s", Number(6))
	}
}

func TestHybrid(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("Hybrid: %s", Hybrid(6))
	}
}
