package goreloaded

import "strings"

func FixSingleQuotes(s string) string {
	runes := []rune(s)
	res := []rune{}
	open := -1

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if r == '\'' {
			if i > 0 && i+1 < len(runes) && IsAlpha(runes[i-1]) && IsAlpha(runes[i+1]) {
				res = append(res, r)
				continue
			}

			if open == -1 {
				open = len(res)
				if len(res) > 0 && res[len(res)-1] != ' ' && res[len(res)-1] != '\n' && res[len(res)-1] != '\t' {
					if chechCloseSingleQuote(runes[i+1:]) {
						res = append(res, ' ')
						open += 1
					} else {
						open = -1
					}
				}
				res = append(res, '\'')
				continue
			}

			if open != -1 {
				content := strings.TrimSpace(string(res[open+1:]))

				res = res[:open+1]

				res = append(res, []rune(content)...)

				res = append(res, '\'')

				if i+1 < len(runes) && runes[i+1] != ' ' && !Checksymbols(string(runes[i+1])) {
					res = append(res, ' ')
				}

				open = -1
				continue
			}
		} else {
			res = append(res, r)
		}
	}

	return string(res)
}

func chechCloseSingleQuote(runes []rune) bool {
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\'' {
			if i > 0 && i+1 < len(runes) && IsAlpha(runes[i-1]) && IsAlpha(runes[i+1]) {
				continue
			}
			return true
		}
	}
	return false
}
