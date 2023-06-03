package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	string := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(string))
}
