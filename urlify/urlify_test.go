package urlify

import (
	"reflect"
	"testing"
)

func TestUrlify(t *testing.T) {
	var tests = []struct {
		s     []byte
		ahead int
		want  []byte
	}{
		{
			s:     []byte("hello world  "),
			ahead: 10,
			want:  []byte("hello%20world"),
		},
		{
			s:     []byte("hello world by me      "),
			ahead: 16,
			want:  []byte("hello%20world%20by%20me"),
		},
	}
	for _, test := range tests {
		urlify(test.s, test.ahead)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("TestUrlify. Got = %s want = %s", test.s, test.want)

		}
	}
}