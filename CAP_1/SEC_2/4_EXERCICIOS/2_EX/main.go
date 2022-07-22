package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Print(idx)
		fmt.Print(": ")
		fmt.Println(arg)
	}
}
