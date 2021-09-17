package equilibrium

import "testing"

func TestEquillibrium(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{
			input: []int{-7, 1, 5, 2, -4, 3, 0},
			want:  3,
		},
		{
			input: []int{-7, 1, 5},
			want:  -1,
		},
	}

	for _, test := range tests{
		got := equilibriumIndex(test.input)
		if got != test.want{
			t.Errorf("TestEquillibrium failed. Got = %d Want = %d",got, test.want)
		}
	}

}
