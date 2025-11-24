package goreloaded

import (
	"strconv"
	"strings"
	"unicode"
)

func ReBuildText(s string) string {
	runes := []rune(s)
	res := ""
	buffer := ""
	inParen := false
	inExit := false
	for l, r := range runes {
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
			outSpace := false
			if l+1 < len(runes) && runes[l+1] == ' ' {
				outSpace = true
			}
			res = newstring("("+buffer+")", res, outSpace)
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

func newstring(style, res string, outSpace bool) string {
	arr := strings.Split(style[1:len(style)-1], ", ")
	resualt := ""
	if len(arr) == 1 {
		arr = append(arr, "1")
	}
	if len(arr) == 2 {
		switch arr[0] {
		case "cap":
			res, allwords, boole := newWord(res, arr[1])
			if !boole {
				resualt = res + style
			}
			return res + Capitalize(allwords)
		case "low":
			res, allwords, boole := newWord(res, arr[1])
			if !boole {
				resualt = res + style
			}
			for _, v := range allwords {
				res += string(unicode.ToLower(v))
			}
			return res
		case "up":
			res, allwords, boole := newWord(res, arr[1])
			if !boole {
				resualt = res + style
			}
			for _, v := range allwords {
				res += string(unicode.ToUpper(v))
			}
			return res
		case "bin":
			if style == "(bin)" {
				return HexAndBiniryWord(res, arr[1], style, 2)
			}
			resualt = res + style
		case "hex":
			if style == "(hex)" {
				return HexAndBiniryWord(res, arr[1], style, 16)
			}
			resualt = res + style
		default:
			resualt = res + style
		}
	}
	resualt = res + style
	if outSpace {
		resualt += " "
	}
	return resualt
}

func newWord(res, nbrword string) (string, string, bool) {
	n, err := strconv.Atoi(nbrword)
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

func CheckWord(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, i := range s {
		if unicode.IsLetter(i)|| unicode.IsDigit(i) {
			return true
		}
	}
	return false
}

func Capitalize(word string) string {
    firstchar := true
    res := ""
    for _, r := range word {
        if r == ' ' {
            res += string(r)
            firstchar = true
            continue
        }

        if (unicode.IsLetter(r) || unicode.IsDigit(r)) && firstchar {
            res += string(unicode.ToUpper(r))
            firstchar = false
        } else {
            res += string(unicode.ToLower(r))
        }
    }
    return res
}

func ConvertToDicemal(nbr string, x int) (int64, error) {
	return strconv.ParseInt(nbr, x, 64)
}

func TrimSpacEendOne(res string) string {
	runes := []rune(res)
	if runes[len(runes)-1] == ' ' {
		return string(runes[:len(runes)-1])
	}
	return res
}

func HexAndBiniryWord(res, nbr, style string, x int) string {
	res, allwords, boole := newWord(res, nbr)
	if !boole {
		return res + style + " "
	}
	arrs := strings.Split(allwords, " ")
	for i := 0; i < len(arrs); i++ {
		if arrs[i] != "" {
			bin, err := ConvertToDicemal(arrs[i], x)
			if err != nil {
				continue
			}
			arrs[i] = strconv.Itoa(int(bin))
		}
	}
	allwords = strings.Join(arrs, " ")
	return res + allwords
}
