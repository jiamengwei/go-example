package main

import (
	"fmt"
	"os"
)

func main() {
	str, _ := os.Getwd()
	fmt.Println(str)
	fmt.Printf("%0.1s", "123456")
}
