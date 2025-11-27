package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	goreloaded "goreloaded/Dependencies"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error : enter 'go run main.go FileInput.txt FileOutput.txt'")
		return
	}
	if !NameFile(strings.TrimSpace(os.Args[1])) || !NameFile(strings.TrimSpace(os.Args[2])) {
		fmt.Println("Error : enter 'go run main.go FileInput.txt FileOutput.txt' name finish by .txt")
		return
	}
	file, err := os.Open(os.Args[1]) // fd == file descriptor
	// buffer read in golang
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	context, err := io.ReadAll(file) // EOF end of file == offset in file == compare ofsset with the file length
	if err != nil {
		fmt.Println(err)
		// log.Println(err) print to stderr
		return
	}
	context = StartByNewLine(string(context))
	newfile, newerr := os.OpenFile(os.Args[2],os.O_CREATE| os.O_WRONLY|os.O_TRUNC, 0o000) // permission // trunc append
	if newerr != nil {
		// fmt.Println(newerr)
		log.Println(newerr)
		// os.Exit(1)
		return
	}
	defer newfile.Close() // leak of resources specially here file descriptor
	_, err = newfile.Write(context)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func StartByNewLine(s string) []byte {
	slice := strings.Split(s, "\n")
	for i := 0; i < len(slice); i++ {
		slice[i] = goreloaded.ReBuildText(string(slice[i]))
		slice[i] = goreloaded.Cleantext(string(slice[i]))
		slice[i] = goreloaded.FixSingleQuotes(string(slice[i]))
		slice[i] = goreloaded.AToAn(string(slice[i]))
	}
	s = strings.Join(slice, "\n")
	return []byte(s)
}

func NameFile(s string) bool {
	if len(s) <= 4 {
		return false
	}
	return strings.HasSuffix(s, ".txt")
}
