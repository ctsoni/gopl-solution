package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("Index %d: ", i)
		fmt.Println(os.Args[i])
	}
}
