package rotatedstring

import "testing"

func TestRotatedString(t *testing.T){
	var tests = []struct{
		input string
		sub string
		want bool
	}{
		{
			input: "waterbottle",
			sub:"bottlewater",
			want: true,
		},
		{
			input: "waterbottle",
			sub:"bottledwater",
		},
	}
	for _, test := range tests{
		got := isRotated(test.input, test.sub)
		if got != test.want{
			t.Errorf("TestRotatedString: got = %v want = %v",got,test.want)
		}
	}
}