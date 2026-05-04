package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: grepgo <what you wanna search>")
	}
	what := os.Args[1]
	fmt.Printf("searching for %s...\n", what)

}
