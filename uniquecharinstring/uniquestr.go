package uniquestring

// isUnique1 use bruteforce to check if chars are unique
// O(n^2)
func isUnique1(s string)bool{
	for i := range s {
		for j:=0 ;j < len(s); j++{
			if i==j {
				continue
			}
			if s[i] == s[j]{
				return false
			}
		}
	}
	return true
}
// O(n) with O(n) extra space
func isUnique2(s string)bool{
	seen := make(map[byte]bool)

	for i := range s {
		if ok:= seen[s[i]]; ok {
			return false
		}
		seen[s[i]] = true
	}
	return true
}
// O(n) with O(1) extra space
func isUnique3(s string)bool{

	firstHalf := uint64(0) // 0- 63
	secondHalf := uint64(0) // 64 -127

	for i := range s {
		switch{
		case int(s[i]) < 64:
			if firstHalf & (1 << int(s[i])) > 0{
				return false
			}
			firstHalf |= uint64(1 << int(s[i]))
		default:
			if secondHalf & (1 << (int(s[i] - 64))) > 0{
				return false
			}
			secondHalf |= uint64(1 << (int(s[i] - 64)))
		}
	}
	return true
}