package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 5)
	}
}
