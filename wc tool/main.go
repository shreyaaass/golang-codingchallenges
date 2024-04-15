package main

import (
	"fmt"
	"os"
	"flag"
	"unicode/utf8"
	"strings"
)

func main() {
	enableBytes := flag.Bool("c", false, "outputs the number of bytes in a file")
	enableLines := flag.Bool("l", false, "outputs the number of lines in a file")
	enableWords := flag.Bool("w", false, "outputs the number of words in a file")
	enableChars := flag.Bool("m", false, "outputs the number of characters in a file")
	file_path := flag.String("f", "test.txt", "path to the file")
	var bytesOrNot int=2
	flag.Parse()
	data, err := os.ReadFile(*file_path)
	if err != nil {
		fmt.Print("Error reading file:", err)
		return
	}
	if flag.NFlag()==0{
		*enableBytes=true
		*enableLines=true
		*enableWords=true
	}
	if *enableBytes{
		bytesOrNot=0
	}
	if *enableChars{
		bytesOrNot=1
	}
	if bytesOrNot==0{
		fmt.Print(len(data), "	")
	}else if bytesOrNot==1{
		fmt.Print(utf8.RuneCountInString(string(data)),"	")
	}
	if *enableLines{
		fmt.Print(strings.Count(string(data), string("\n")),"	")
	}
	if *enableWords{
		fmt.Print(strings.Count(string(data),string(" ")), "	")
	}
	fmt.Println(*file_path)
}