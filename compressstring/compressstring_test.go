package compressstring

import "testing"

func TestCompressString(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{
			input:  "aabbccdd",
			output: "aabbccdd",
		},
		{
			input:  "aaaabb",
			output: "a4b2",
		},
	}
	for _, test := range tests {
		got := compress(test.input)
		if got != test.output {
			t.Errorf("TestCompressString failed. Got = %s Want =%s", got, test.output)
		}
	}
}