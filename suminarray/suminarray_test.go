package suminarray

import "testing"

func TestSumInArray(t *testing.T) {
	var tests = []struct {
		input []int
		sum   int
		want  bool
	}{
		{
			input: []int{10, 15, 3, 7},
			sum:   17,
			want:  true,
		},
		{
			input: []int{10, 4, 4, 7},
			sum:   8,
			want:  true,
		},
		{
			input: []int{10, 15, 3, 7},
			sum:   20,
		},
	}
	for _, test := range tests {
		got:= suminarray(test.input, test.sum)
		if got != test.want {
			t.Errorf("TestSumInArray failed. Got = %v Want = %v", got, test.want)
		}
	}
}

func TestSumInSortedArray(t *testing.T) {
	var tests = []struct {
		input []int
		sum   int
		want  bool
	}{
		{
			input: []int{3,7,10, 15},
			sum:   17,
			want:  true,
		},
		{
			input: []int{4, 4, 7, 10},
			sum:   8,
			want:  true,
		},
		{
			input: []int{3, 7, 15, 20},
			sum:   20,
		},
	}
	for _, test := range tests {
		got:= suminarray_sorted(test.input, test.sum)
		if got != test.want {
			t.Errorf("TestSumInSortedArray failed. Got = %v Want = %v", got, test.want)
		}
	}
}

func TestCountpairswithsum(t *testing.T) {
	var tests = []struct {
		input []int
		sum   int
		want  int
	}{
		{
			input: []int{3,7,10, 15},
			sum:   17,
			want:  1,
		},
		{
			input: []int{3,3,3,3,3,9,-3},
			sum:   6,
			want:  5,
		},
		{
			input: []int{3, 7, 15, 20},
			sum:   20,
		},
	}
	for _, test := range tests {
		got:= countpairswithsum(test.input, test.sum)
		if got != test.want {
			t.Errorf("TestCountpairswithsum failed. Got = %d Want = %d", got, test.want)
		}
	}
}