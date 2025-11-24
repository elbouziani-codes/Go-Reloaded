package goreloaded

func Cleantext(s string) string {
	runes := []rune(s)
	res := ""
	n := len(runes)
	for i := 0; i < n; i++ {
		if Checksymbols(string(runes[i])) {
			res = TrimSpaceWhiteEnd(res)
			for i < n && Checksymbols(string(runes[i])) {
				res += string(runes[i])
				i++
			}
			if i >= n {
				break
			}
			for i < n && runes[i] == ' ' {
				i++
			}
			if i >= n {
				break
			}
			if Checksymbols(string(runes[i])) {
				i -= 1
				continue
			}
			if i < n {
				res += " " + string(runes[i])
			}
		} else {
			res += string(runes[i])
		}
	}
	return res
}

func TrimSpaceWhiteEnd(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	i := len(runes) - 1
	for i >= 0 && (runes[i] == ' ' || runes[i] == '\n' || runes[i] == '\t') {
		i--
	}
	return string(runes[:i+1])
}