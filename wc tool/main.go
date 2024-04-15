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
	enableLines := flag.Bool("l", false, "outputs the number of bytes in a file")
	enableWords := flag.Bool("w", false, "outputs the number of bytes in a file")
	enableChars := flag.Bool("m", false, "outputs the number of characters in a file")
	file_path := flag.String("f", "test.txt", "path to the file")

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
		fmt.Print(len(data),"	")
	}
	if *enableLines{
		fmt.Print(strings.Count(string(data), string("\n")),"	")
	}
	if *enableWords{
		fmt.Print(strings.Count(string(data),string(" ")), "	")
	}
	if *enableChars{
		fmt.Print(utf8.RuneCountInString(string(data)),"	")
	}
	fmt.Println(*file_path)

}



// package main

// import (
// 	"flag"
// 	"fmt"
// 	"io/ioutil"
// 	"strings"
// 	"unicode/utf8"
// )

// func count(data []byte, separator byte) int {
// 	return strings.Count(string(data), string(separator))
// }

// func main() {
// 	file_path:=flag.String("f", "test.txt", "path to the file")
// 	enableBytes := flag.Bool("c", true, "outputs the number of bytes in a file")
// 	enableLines := flag.Bool("l", true, "outputs the number of lines in a file")
// 	enableWords := flag.Bool("w", true, "outputs the number of words in a file")
// 	enableChars := flag.Bool("m", false, "outputs the number of characters in a file")

// 	data, err := ioutil.ReadFile(*file_path)
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// 	flag.Parse()
	
// 	if *enableBytes {
// 		fmt.Println("Bytes:", len(data))
// 	}
// 	if *enableLines {
// 		fmt.Println("Lines:", count(data, '\n')+1)
// 	}
// 	if *enableWords {
// 		words := strings.Fields(string(data))
// 		fmt.Println("Words:", len(words))
// 	}
// 	if *enableChars {
// 		fmt.Println("Characters:", utf8.RuneCountInString(string(data)))
// 	}
// }
