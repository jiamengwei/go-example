package main

import "fmt"

func main() {
	//声明一个长度为2的数组
	var arrA [2]int
	//输出数组的默认值
	fmt.Println(arrA) //[0 0]
	//为数组赋值
	arrA[0] = 1
	arrA[1] = 2
	fmt.Println(arrA) //[1 2]

	arrB := [3]int{1, 2, 3}
	fmt.Println(arrB) //[1 2 3]

	arrC := [3]int{1}
	fmt.Println(arrC) //[1 0 0]

	arrD := [...]int{1, 2, 3, 4}
	fmt.Println(arrD) //[1 2 3 4]

	arrE := [2]int{1, 2}
	//arrF := [3]int{}  长度不同的数组被视为两种类型，无法赋值
	arrF := [2]int{}
	arrE = arrF
	fmt.Println(arrE) //[0 0]

	arrG := [...]string{"A", "B", "C"}
	arrH := arrG
	arrH[0] = "Z"
	fmt.Println(arrG) //[A B C]
	fmt.Println(arrH) //[Z B C]

	arrI := [2]int{1,2}
	arrJ := changeArr(arrI)
	fmt.Println(arrI)	//[1 2]
	fmt.Println(arrJ)	//[99 2]

	arrK := [...]int{1,2,3,4}
	fmt.Println(len(arrK))	//4

	for i:=0; i<len(arrK); i++ {
		fmt.Printf("%d \n",arrK[i])
	}

	for i, v := range arrK{
		fmt.Printf("i:%d, v:%d\n", i, v)
	}

	var arrL [3][3]string
	arrL[0][0] = "A"
	arrL[0][1] = "B"
	arrL[0][2] = "C"

	arrL[1][0] = "D"
	arrL[1][1] = "E"
	arrL[1][2] = "F"

	arrL[2][0] = "G"
	arrL[2][1] = "H"
	arrL[2][2] = "I"

	printMulitArr(arrL)
}

func printMulitArr(arr [3][3]string){
	for _, v1 := range arr{
		for _, v2 := range v1{
			fmt.Println(v2)
		}
	}
}

func changeArr(arr [2]int) [2]int {
	arr[0] = 99
	fmt.Println(arr)
	return arr
}
