package equilibrium

// equilibriumIndex returns equillibrium
// index of the input inputay.
func equilibriumIndex(input []int) int {
	sright := 0
	for i := 0; i < len(input); i++ {
		sright += input[i]
	}
	sleft := 0
	for i := 0; i < len(input); i++ {
		sright -= input[i]
		if sleft == sright {
			return i
		}
		sleft += input[i]
	}
	return -1
}
