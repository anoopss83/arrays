package permutnsubstr


func isPermutation(str1, str2 string)bool{

	if len(str2) != len(str1) {
		return false
	}
	counts := make(map[byte]int)
	// Populate count
	for i := range str1{
		counts[str1[i]] += 1
	}

	for i := range str2 {
		counts[str2[i]] -= 1
		if counts[str2[i]] <0 {
			return false
		}
	}
	return true

}