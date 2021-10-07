package maxactivities

import (
	"reflect"
	"testing"
)

func TestMaxActivities(t *testing.T) {
	var tests = []struct {
		intervals  []interval
		activities []int
	}{
		{
			intervals: []interval{
				{start: 1, end: 2},
				{start: 3, end: 4},
				{start: 0, end: 6},
				{start: 5, end: 7},
				{start: 8, end: 9},
				{start: 5, end: 9},
			},
			activities: []int{0, 1, 3, 4},
		},
	}

	for _, test := range tests {
		sortIntervals(test.intervals)

		want := maxActivities(test.intervals)
		if !reflect.DeepEqual(want, test.activities) {
			t.Errorf("TestMaxActivities failed. Got = %v Want = %v", test.activities, want)
		}
	}
}
