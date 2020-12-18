package main

import "fmt"

/**
指针是一个变量，它存储另一个变量的内存地址。
*T 表示T类型的指针
& 取一个变量的地址
*/
func main() {
	strA := "A"
	//指针adrA存储了变量strA的内存地址
	var adrA *string
	if adrA == nil {
		fmt.Println("adrA is nil")
	}

	adrA = &strA
	fmt.Println("strA的地址", adrA)
	fmt.Printf("adrA的类型：%T \n", adrA)

	size := new(int)
	fmt.Printf("type: %T, value : %d, address: %v \n", size, *size, size)
	*size = 23
	fmt.Printf("type: %T, value : %d, address: %v \n", size, *size, size)

	intA := 233
	adrB := &intA
	*adrB++
	fmt.Println(intA) //234
	change(adrB)
	fmt.Println(intA) //100

	arrA := [3]int{1, 2, 3}
	modify(&arrA)
	fmt.Println(arrA)

	modify2(arrA[:])
	fmt.Println(arrA)
}

func change(i *int) {
	*i = 100
}

func modify(arr *[3]int) {
	//(*arr)[0] = 90
	arr[0] = 90
}

func modify2(sls []int){
	sls[0] = 99
}
