package main

import "fmt"

/**
defer用于在函数返回之前执行指定的调用
defer 方法在接受参数时，参数的值在程序执行到defer语句时已经确定，而不是在正真执行时确定参数的值

*/
func hello() {
	defer fmt.Println("world !")
	fmt.Printf("hello ")
}

func helloPeople(people string) {
	defer fmt.Println("hello ", people)
}

func helloNumber() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("hello ", i)
	}
}

func main() {
	hello()
	people := "Jack"
	defer helloPeople(people)
	people = "John"
	fmt.Println("people is ", people)
	helloNumber()
}
