package main

import "fmt"

func main() {
	var a []int
	a = append(a, 2)
	a = append(a, 1)
	fmt.Println(a)

	println(len(a))
	println(cap(a))
	var s [23]int
	println(s)

}
