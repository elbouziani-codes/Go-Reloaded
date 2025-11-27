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
	isOpen := -1
	SkipSpace := false
	for l, r := range runes {
		if SkipSpace {
			SkipSpace = false
			if r == ' ' {
				continue
			}
		}
		if r == '(' {
			isOpen = l
			if isOpen == 0 || (isOpen > 0 && runes[isOpen-1] == ' ') {
				if inParen {
					res += "(" + buffer
				}
				inParen = true
				buffer = ""
				continue
			}
		}
		if r == ')' && inParen {			
			if l == len(runes)-1 || (l < len(runes)-1 && runes[l+1] == ' ') {
				inParen = false
				res = newstring("("+buffer+")", res, isOpen == 0,l == len(runes)-1)
				buffer = ""
				isOpen = -1
				SkipSpace = true
				continue
			}
		}
		if inParen {
			buffer += string(r)
		} else {
			res += string(r)
		}
	}
	if buffer != "" {
		res += "(" + buffer
	}
	return res
}

func newstring(style, res string, first,last bool) string {
	arr := strings.Split(style[1:len(style)-1], ", ")
	result := ""
	if len(arr) == 1 {
		arr = append(arr, "1")
	}
	if len(arr) == 2 {
		switch arr[0] {
		case "cap":
			res, allwords, boole := newWord(res, arr[1], style)
			if !boole {
				break
			}
			return res + Capitalize(allwords)
		case "low":
			res, allwords, boole := newWord(res, arr[1], style)
			if !boole {
				break
			}
			for _, v := range allwords {
				res += string(unicode.ToLower(v))
			}
			return res
		case "up":
			res, allwords, boole := newWord(res, arr[1], style)
			if !boole {
				break
			}
			for _, v := range allwords {
				res += string(unicode.ToUpper(v))
			}
			return res
		case "bin":
			if style == "(bin)" {
				return HexAndBiniryWord(res, arr[1], style, 2)
			}
		case "hex":
			if style == "(hex)" {
				return HexAndBiniryWord(res, arr[1], style, 16)
			}
		}
	}
	result = res +style
	if !last {
		result += " " 
	}
	return result
}

func newWord(res, nbrword, style string) (string, string, bool) {
	n, err := strconv.Atoi(nbrword)
	if err != nil {
		if strings.HasSuffix(err.Error(), "value out of range") {
			return res, "", false
		}
		return res, style, false
	} else if n <= 0 {
		return res, "", true
	}
	first := false
	words := strings.Split(res," ")
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
		if unicode.IsLetter(i) || unicode.IsDigit(i) {
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
	if len(res) == 0 {
		return res
	}
	runes := []rune(res)
	if runes[len(runes)-1] == ' ' {
		return string(runes[:len(runes)-1])
	}
	return res
}

func HexAndBiniryWord(res, nbr, style string, x int) string {
	res, allwords, boole := newWord(res, nbr, style)
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
