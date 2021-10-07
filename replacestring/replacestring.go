package replacestring

func compare1(s1,s2 string)bool{

	// find if replace OR insert/removal
	diff := len(s1) - len(s2)
	if diff > 1 || diff < -1 {
		return false
	}

	if diff == 0 {
		return replace(s1,s2)
	}
	
	return insertORRemove(s1,s2)

}

func insertORRemove(s1,s2 string)bool {
	i:=0
	j := 0
	for i < len(s1) && j < len(s2){
		if s1[i] != s2[j]{
			if i != j {
				return false
			}
			
			if len(s1) > len(s2){
				i++
			}else{
				j++
			}
			continue
		}
		i++
		j++
	}
	return true
}
func replace(s1, s2 string)bool{
	foundDiff := false
	for i:=0 ; i < len(s1); i++ {
		if s1[i] != s2[i]{
			if foundDiff{
				return false
			}
			foundDiff = true
		}
	}
	return true
}

func compare(s1, s2 string) bool {

	diff := len(s1) - len(s2)
	if diff > 1 || diff < -1 {
		return false
	}
	m := make(map[byte]int)

	if len(s1) < len(s2) {
		insertToMap(s1, m)
		return checkMap(s2, m)
	}
	insertToMap(s2, m)
	return checkMap(s1, m)

}
func checkMap(s string, m map[byte]int) bool {
	diff := 1
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			diff--
		}
		if diff < 0{
			return false
		}
		v := m[s[i]]
		if v == 0 {
			delete(m,s[i])
			continue
		}
		m[s[i]] = v - 1
	}
	return true

}
func insertToMap(s string, m map[byte]int) {
	for i := range s {
		m[s[i]] = m[s[i]] + 1
	}
}
