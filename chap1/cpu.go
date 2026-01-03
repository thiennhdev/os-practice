package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: cpu <string>\n")
		os.Exit(1)
	}

	str := os.Args[1]

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
