package uniquestring

import "testing"

func Test_isUnique(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{
			input: "abcdef",
			want:  true,
		},
		{
			input: "aaa",
		},
	}

	for _, test := range tests {
		got := isUnique1(test.input)
		if got != test.want {
			t.Errorf("Test_isUnique1 failed. Got = %v Want = %v",got, test.want)
		}
		got = isUnique2(test.input)
		if got != test.want {
			t.Errorf("Test_isUnique2 failed. Got = %v Want = %v",got, test.want)
		}
		got = isUnique3(test.input)
		if got != test.want {
			t.Errorf("Test_isUnique3 failed. Got = %v Want = %v input = %v",got, test.want, test.input)
		}
	}

}