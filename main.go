package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Pass source file as argument")
		return
	}
	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	tokens, err := Scan(source)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range tokens {
		fmt.Println(t)
	}
}
