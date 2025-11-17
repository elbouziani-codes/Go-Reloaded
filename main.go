package main

import (
	"fmt"
	"io"
	"os"
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
	newfile , newerr := os.OpenFile(os.Args[2],os.O_CREATE&os.O_WRONLY&os.O_TRUNC,0777)
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