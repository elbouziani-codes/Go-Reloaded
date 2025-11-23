package goreloaded

import (
	"strconv"
	"strings"
)

func CheckWord(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || (i >= '0' && i <= '9') {
			return true
		}
	}
	return false
}

func Capitalize(s string) string {
	firstword := true
	res := ""
	for _, i := range s {
		if firstword {
			if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || (i >= '0' && i <= '9') {
				if i >= 'a' && i <= 'z' {
					res += strings.ToUpper(string(i))
				} else {
					res += string(i)
				}
				firstword = false
			} else {
				res += string(i)
				continue
			}
		} else {
			if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || (i >= '0' && i <= '9') {
				if i >= 'A' && i <= 'Z' {
					res += strings.ToLower(string(i))
				} else {
					res += string(i)
				}
			} else {
				res += string(i)
				firstword = true
				continue
			}
		}
	}
	return res
}

func ConvertToDicemal(nbr string, x int) (int64, error) {
	return strconv.ParseInt(nbr, x, 64)
}

func TrimSpaceEnd(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	i := len(runes) - 1
	for i >= 0 && runes[i] == ' ' {
		i--
	}
	return string(runes[:i+1])
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

func TrimSpacEendOne(res string) string {
	runes := []rune(res)
	if runes[len(runes)-1] == ' ' {
		return string(runes[:len(runes)-1])
	}
	return res
}

func Checksymbols(s string) bool {
	runes := []rune{'.', ',', '!', '?', ':', ';'}
	for _, v := range runes {
		if v == rune(s[0]) {
			return true
		}
	}
	return false
}