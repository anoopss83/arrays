package replacestring

import "testing"

func TestReplaceString(t *testing.T) {
	var tests = []struct{
		s1 string
		s2 string
		want bool
	}{
		{
			s1: "pale",
			s2: "pales",
			want: true,
		},
		{
			s2: "pale",
			s1: "pales",
			want: true,
		},
		{
			s2: "pale",
			s1: "bale",
			want: true,
		},
		{
			s2: "bale",
			s1: "bale",
			want: true,
		},
		{
			s2: "pale",
			s1: "bake",
			want: false,
		},
		{
			s2: "apple",
			s1: "aple",
			want: true,
		},
		{
			s2: "pale",
			s1: "ale",
			want: true,
		},
		{
			s2: "pale",
			s1: "ake",
			want: false,
		},
	}
	for _, test := range tests{
		got := compare(test.s1, test.s2)
		if test.want != got {
			t.Errorf("TestReplaceString. Got = %v Want = %v",got, test.want)
		}
		got = compare1(test.s1, test.s2)
		if test.want != got {
			t.Errorf("TestReplaceString. Got = %v Want = %v",got, test.want)
		}
	}
}
