package suminarray


func suminarray(array []int, sum int) bool {
	seen := make(map[int]bool)
	for i := 0; i < len(array); i++ {
		diff := sum - array[i]
		if _, ok := seen[diff]; ok {
			return true
		}
		seen[array[i]] = true

	}
	return false
}

func suminarray_sorted(array []int, sum int) bool {
	lo := 0
	hi := len(array) - 1
	for lo < hi {
		switch {
		case array[lo]+array[hi] == sum:
			return true
		case array[lo]+array[hi] > sum:
			hi--
		case array[lo]+array[hi] < sum:
			lo++
		}
	}
	return false
}

// Elements are not distinct.
func countpairswithsum(array []int, sum int) int {

	seen := make(map[int]struct{})
	count := 0

	for i := 0; i < len(array); i++ {
		diff := sum - array[i]
		if _, ok := seen[diff]; ok {
			count += 1
		}
		seen[array[i]] = struct{}{}
	}
	return count
}
