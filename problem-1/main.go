package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	file_path=os.Args[3]
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(len(data))
}
