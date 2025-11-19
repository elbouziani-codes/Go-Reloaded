package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)
func main(){
	if(len(os.Args) != 3){
		fmt.Println("Error")
		return
	}
	file , err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	context , err := io.ReadAll(file)
	if err != nil{
		fmt.Println(err)
		return
	}
	context = []byte(a(string(context)))
	newfile , newerr := os.OpenFile(os.Args[2],os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0777)
	if newerr != nil{
		fmt.Println(newerr)
		return
	}
	defer newfile.Close()
	_,err = newfile.Write(context)
	if err != nil{
		fmt.Println(err)
		return
	}
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
			if check(arr[j][0]) {
				arr[i] = "an"
			}
		}
	}
	res := strings.Join(arr," ")
	return res
}
func check(s byte) bool{
	var char = []byte{'a','o','u','e','i','A','O','U','E','I'}
	for i := 0; i < len(char); i++ {
		if char[i] == s {
			return true
		}
	}
	return false
}

func symbol