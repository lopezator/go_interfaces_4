package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please, enter the file name as argument!")
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Hey, something went wrong, check if you provided the correct fileName as arg")
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
