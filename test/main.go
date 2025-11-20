package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	context, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	context = []byte(readd(string(context)))
	newfile, newerr := os.OpenFile(os.Args[2], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if newerr != nil {
		fmt.Println(newerr)
		return
	}
	defer newfile.Close()
	_, err = newfile.Write(context)
	if err != nil {
		fmt.Println(err)
		return
	}
}


func readd(s string) string {
	runes := []rune(s)
	res := ""
	buffer := ""
	inParen := false
	inExit := false
	for _, r := range runes {
		if r == '(' {
			if inParen{
				res += "(" + buffer
			}
			inParen = true
			buffer = ""
			continue
		}
		if r == ')' && inParen{
			inParen = false
			res = TrimSpaceend(res)
			res = newstring("(" + buffer + ")", res)
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
			res = TrimSpaceend(res)
			inExit = false
		}
	}
	if buffer != "" {
		res += " (" +buffer
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
	if err != nil || n <= 0 {
		return res, "", false
	}
	words := strings.Split(res, " ")
	count := 0
	for i := len(words) - 1; i >= 0; i-- {
		if checkWord(words[i]) {
			count+=1
			if count == n {
				break
			}
		}
	}
	if count == 0 {
		return res, "", false
	}
	if count < n {
		n = count
	}
	i := len(words)-1
	for ; i >= 0; i-- {
	    if checkWord(words[i]) && count != 0 {
	        count-= 1
	    }
	    if(count == 0){
	        break
	    }
	}
	resBefore := strings.Join(words[i:], " ")
	allwords := strings.Join(words[:i], " ")
	return allwords ,resBefore , true
}

func checkWord(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _,i := range s{
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || (i >= '0' && i <= '9') {
			return true
		}
	}
	return false
}
func Capitalize(s string) string {
	firstword := true
	res := ""
	for _,i := range s{
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

func TrimSpaceend(res string) string{
	runes := []rune(res)
	if runes[len(runes)-1] == ' ' {
		return string(runes[:len(runes)-1])
	}
	return res
	
}
func ConvertToDicemal(nbr string, x int) (int64, error) {
	return strconv.ParseInt(nbr, x, 64)
}
func a(s string)string{
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr)-1; i++ {
		j := i+1
		for ; j < len(arr); j++ {
			if len(arr[j]) > 0 {
				break
			}
		}
		if(j == len(arr)){
			break
		}
		if arr[i] == "a"{
			if check(rune(arr[j][0])) {
				arr[i] = "an"
			}
		}
	}
	res := strings.Join(arr," ")
	return res
}
func check(s rune) bool{
	var char = []rune{'a','o','u','e','i','A','O','U','E','I'}
	for _,i := range char{
		if i == s {
			return true
		}
	}
	return false
}