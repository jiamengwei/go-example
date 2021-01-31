package main

import (
	"fmt"
)

func main() {
	sl := make([]interface{}, 1, 1)
	fmt.Printf("%T", sl)

	a := [2]int{1, 2}
	fmt.Printf("%T", a)

	/**
	通过数组创建一个切片的引用
	*/
	arrA := [...]int{1, 2, 3, 4, 5}
	sliceA := arrA[0:2]
	fmt.Println(sliceA)

	/**
	创建一个返回切片引用的数组
	*/
	sliceB := []int{1, 2, 3, 4}
	fmt.Println(sliceB)

	//遍历一个切片
	for i := range sliceA {
		sliceA[i]++
	}
	fmt.Println(sliceA)
	fmt.Println(arrA)

	//切片本质上是对一个数组的引用，对切片或者数组的修改都会影响彼此
	arrB := [3]int{1, 2, 3}
	sliceC := arrB[0:3]
	sliceD := arrB[:]

	fmt.Print(arrB) //[1 2 3]
	sliceC[0] = 9
	fmt.Print(arrB) //[9 2 3]
	sliceD[1] = 8
	fmt.Println(arrB) //[9 8 3]

	arrB[2] = 7
	fmt.Println(sliceC) //[9 8 7]
	fmt.Println(sliceD) //[9 8 7]

	// len(v Type) 获取切片的长度
	// cap(v Type) 获取切片的容量
	arrC := [4]string{"A", "B", "C", "D"}
	sliceE := arrC[1:2]
	fmt.Println(sliceE)
	fmt.Printf("sliceE len:%d, cap:%d \n", len(sliceE), cap(sliceE)) //sliceE len:1,cap:3

	arrD := [5]string{"A", "B", "C", "D", "E"}
	sliceF := arrD[2:2]
	fmt.Println(sliceF)                                              //[]
	fmt.Printf("sliceF len:%d, cap:%d \n", len(sliceF), cap(sliceF)) //sliceF len:0, cap:3
	sliceF = arrD[1:]
	fmt.Println(sliceF)                                              //[B C D E]
	fmt.Printf("sliceF len:%d, cap:%d \n", len(sliceF), cap(sliceF)) //sliceF len:4, cap:4

	//使用make方法创建一个len为2，cap为4的切片
	sliceG := make([]int, 2, 4)
	fmt.Println(sliceG)                                              //[0 0]
	fmt.Printf("sliceG len:%d, cap:%d \n", len(sliceG), cap(sliceG)) //sliceG len:2, cap:4

	//使用append方法追加元素
	sliceG = append(sliceG, 7, 8, 9)
	fmt.Println(sliceG)                                              //[0 0 7 8 9]
	fmt.Printf("sliceG len:%d, cap:%d \n", len(sliceG), cap(sliceG)) //sliceG len:5, cap:8

	// len和cap都为0的切片等于nil
	var sliceH []int
	if sliceH == nil {
		fmt.Printf("sliceH is empty, len:%d, cap:%d \n", len(sliceH), cap(sliceH)) //sliceH is empty, len:0, cap:0
	}

	//追加元素到nil切片，切片的cap变为追加元素的个数加一
	sliceI := []int{1, 2, 3}
	sliceH = append(sliceH, sliceI...)
	fmt.Printf("sliceH len:%d, cap:%d \n", len(sliceH), cap(sliceH)) //sliceH len:3, cap:4

	//切片是引用类型，在方法中对切片进行修改会影响到当前切片
	sliceJ := []int{1, 2, 3, 4}
	fmt.Println(sliceJ) //[1 2 3 4]
	changeSlice(sliceJ)
	fmt.Println(sliceJ) //[0 1 2 3]

	//创建一个嵌套切片
	sliceK := [][]string{
		{"A", "a"},
		{"B", "b"},
		{"C", "c"},
	}
	//遍历嵌套切片
	printMulitSlice(sliceK)

	//使用copy方法优化内存使用
	arrE := [4]string{"A", "B", "C", "D"}
	sliceL := cutForSlice(arrE, 1, 3)
	fmt.Println(sliceL)
	fmt.Printf("sliceL  len:%d, cap:%d \n", len(sliceL), cap(sliceL))
}

func cutForSlice(arr [4]string, start, end int) []string {
	copySrc := arr[start:end]
	copyDst := make([]string, len(copySrc))
	copy(copyDst, copySrc)
	return copyDst
}

func printMulitSlice(slice [][]string) {
	for _, v1 := range slice {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}

func changeSlice(slice []int) {
	for i := range slice {
		slice[i]--
	}
}
