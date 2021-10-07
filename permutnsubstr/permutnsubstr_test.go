package permutnsubstr

import (
	"testing"
)

func TestPermutnSubStr(t *testing.T){
	var tests = []struct{
		str string
		substr string
		want bool
	}{
		{
			str: "abcab",
			substr:"abcab",
			want:true,
		},
		{
			str: "abba",
			substr:"abc",
			want:false,
		},
			{
			str: "abcd",
			substr:"abce",
			want:false,
		},
	
	}
	for _, test := range tests{
		got := isPermutation(test.str, test.substr)
		if got != test.want {
			t.Errorf("TestPermutnSubStr failed. Got = %v Want = %v", got, test.want)
		}
	}
}
