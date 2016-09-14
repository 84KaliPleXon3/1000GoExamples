package main

import "testing"

func TestFoo1(t *testing.T) {
	var tests = []struct {
		inputM int
		inputN int
		want   string
	}{
		{1, 5, "15"},
	}
	for _, test := range tests {
		if got := foo1(test.inputM, test.inputN); got != test.want {
			t.Errorf("foo1(%v, %v) = %v", test.inputM, test.inputN, got)
		}
	}
}
