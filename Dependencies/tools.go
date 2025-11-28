package goreloaded

func Checksymbols(s string) bool {
	runes := []rune{'.', ',', '!', '?', ':', ';'}
	for _, v := range runes {
		if v == rune(s[0]) {
			return true
		}
	}
	return false
}

func IsAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
