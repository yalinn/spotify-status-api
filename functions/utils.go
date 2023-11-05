package functions

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

func Cryptit(key string, stc bool) string {
	runes := []rune(key)
	i := 0
	for _, r := range runes {
		if stc {
			runes[i] = r + 10
		} else {
			runes[i] = r - 10
		}
		i++
		stc = !stc
	}
	return string(runes)
}
