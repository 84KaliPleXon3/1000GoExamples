package main

import "testing"

func TestLoop(t *testing.T) {
	var tests = []struct {
		inputM int
		inputN int
		want   []int
	}{
		{1, 2, []int{1, 2}},
	}
	for _, test := range tests {
		if got := loop(test.inputM, test.inputN); equal(got, test.want) {
			t.Errorf("loop(%v, %v) = %v", test.inputM, test.inputN, got)
		}
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
