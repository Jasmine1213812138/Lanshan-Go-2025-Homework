package main

import "testing"

func TestCalculate(t *testing.T) {
	test := []struct {
		name string
		got  int
		want int
	}{
		{"add", add(3, 6), 9},
		{"subtract", subtract(8, 9), -1},
		{"mul", mul(3, 3), 9},
		{"div", div(8, 2), 4},
	}
	for _, tes := range test {
		if tes.want != tes.got {
			t.Errorf("%s = %d, want %d", tes.name, tes.got, tes.want)
		}
	}
}
