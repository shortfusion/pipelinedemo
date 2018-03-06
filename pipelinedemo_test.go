package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestDifference(t *testing.T) {
	diff := Difference(10, 5)
	if diff != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", diff, 5)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(2*i, 2+i)
	}
}
