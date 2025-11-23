package goreloaded

import (
	"strconv"
	"strings"
)

func ReBuildText(s string) string {
	runes := []rune(s)
	res := ""
	buffer := ""
	inParen := false
	inExit := false
	for _, r := range runes {
		if r == '(' {
			if inParen {
				res += "(" + buffer
			}
			inParen = true
			buffer = ""
			continue
		}
		if r == ')' && inParen {
			inParen = false
			res = newstring("("+buffer+")", res)
			inExit = true
			buffer = ""
			continue
		}

		if inParen {
			buffer += string(r)
		} else {
			res += string(r)
		}
		if inExit {
			res = TrimSpacEendOne(res)
			inExit = false
		}
	}
	if buffer != "" {
		res += " (" + buffer
	}
	return res
}

func newstring(style, res string) string {
	arr := strings.Split(style[1:len(style)-1], ", ")

	if len(arr) == 1 {
		arr = append(arr, "1")
	} else if len(arr) > 2 {
		return res + style + " "
	}
	switch arr[0] {
	case "cap":
		res, allwords, boole := newWord(res, arr[1])
		if !boole {
			return res + style + " "
		}
		return res + Capitalize(allwords)
	case "low":
		res, allwords, boole := newWord(res, arr[1])
		if !boole {
			return res + style + " "
		}
		return res + strings.ToLower(allwords)
	case "up":
		res, allwords, boole := newWord(res, arr[1])
		if !boole {
			return res + style + " "
		}
		return res + strings.ToUpper(allwords)
	case "bin":
		res, allwords, boole := newWord(res, arr[1])
		if !boole {
			return res + style + " "
		}
		arrs := strings.Split(allwords, " ")
		for i := 0; i < len(arrs); i++ {
			if arrs[i] != "" {
				bin, err := ConvertToDicemal(arrs[i], 2)
				if err != nil {
					continue
				}
				arrs[i] = strconv.Itoa(int(bin))
			}
		}
		allwords = strings.Join(arrs, " ")
		return res + allwords
	case "hex":
		res, allwords, boole := newWord(res, arr[1])
		if !boole {
			return res + style + " "
		}
		arrs := strings.Split(allwords, " ")
		for i := 0; i < len(arrs); i++ {
			if arrs[i] != "" {
				bin, err := ConvertToDicemal(arrs[i], 16)
				if err != nil {
					continue
				}
				arrs[i] = strconv.Itoa(int(bin))
			}
		}
		allwords = strings.Join(arrs, " ")
		return res + allwords
	default:
		return res + style + " "
	}
}

func newWord(res, nbrword string) (string, string, bool) {
	n, err := strconv.Atoi(strings.TrimSpace(nbrword))
	if err != nil {
		return res, "", false
	} else if n <= 0 {
		return res, "", true
	}
	first := false
	words := strings.Split(res, " ")
	count := 0
	for i := len(words) - 1; i >= 0; i-- {
		if CheckWord(words[i]) {
			count += 1
			if i == 0 {
				first = true
			}
			if count == n {
				break
			}
		}
	}
	if count == 0 {
		return res, "", true
	}
	if count < n {
		n = count
	}
	i := len(words) - 1
	for ; i >= 0; i-- {
		if CheckWord(words[i]) && count != 0 {
			count -= 1
		}
		if count == 0 {
			break
		}
	}
	allwords := strings.Join(words[:i], " ")
	resBefore := strings.Join(words[i:], " ")
	if first {
		return allwords, resBefore, true
	}
	return allwords, " " + resBefore, true
}
