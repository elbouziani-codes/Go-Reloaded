package goreloaded

import (
	"strings"
)

func AToAn(s string) string {
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr)-1; i++ {
		j := i + 1
		for ; j < len(arr); j++ {
			if len(arr[j]) > 0 {
				break
			}
		}
		if j == len(arr) {
			break
		}
		if len(arr[i]) != 0 {
			a, bools := Check_A(arr[i])
			if bools {
				if Check(rune(arr[j][0])) {
					a += "n"
				}
			}
			arr[i] = a
		}
	}
	res := strings.Join(arr, " ")
	return res
}

func Check_A(s string) (string, bool) {
	runes := []rune(s)
	if !strings.HasSuffix(s, "a") {
		return s, false
	}
	for i := len(runes) - 2; i >= 0; i-- {
		if i < 0 {
			break
		}
		if IsAlpha(runes[i]) || (runes[i] >= '0' && runes[i] <= '9') {
			return s, false
		}
	}
	return s, true
}

func Check(s rune) bool {
	char := []rune{'a', 'o', 'u', 'e', 'i', 'h', 'H', 'A', 'O', 'U', 'E', 'I'}
	for _, i := range char {
		if i == s {
			return true
		}
	}
	return false
}
