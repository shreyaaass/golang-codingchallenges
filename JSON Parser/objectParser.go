package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("object.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

}
