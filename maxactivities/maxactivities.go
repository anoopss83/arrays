package maxactivities

import "fmt"

type interval struct {
	start int
	end   int
}

func maxActivities(intervals []interval) []int {
	var activities []int
	activities = append(activities, 0)
	lastEnd := 0
	for i := 1; i < len(intervals); i++ {
		if len(activities) != 0 {
			lastEnd = intervals[activities[len(activities) - 1]].end
		}
		if intervals[i].start > lastEnd{
			activities = append(activities, i)
		}
	}

	return activities
}

func sortIntervals(input []interval) {
	var curr interval
	for i := 1; i < len(input); i++ {
		curr = input[i]
		j := i
		for ; j > 0 && curr.end < input[j].end; j-- {
			input[j].end = input[j-1].end
			input[j].start = input[j-1].start
		}
		input[j].end = curr.end
		input[j].start = curr.start
	}
	fmt.Println("Sorted interval ", input)
}
