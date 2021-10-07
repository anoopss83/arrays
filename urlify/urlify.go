package urlify

func urlify(s []byte, ahead int) {

	nSpaces := (len(s) - ahead) / 2
	tail := len(s) - 1

	for nSpaces > 0 {
		switch {
		case s[ahead] == ' ':
			s[tail] = '0'
			s[tail-1] = '2'
			s[tail-2] = '%'
			tail = tail - 2
			nSpaces--
		default:
			s[tail] = s[ahead]
		}
		tail--
		ahead--
	}

}