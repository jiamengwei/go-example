package main

import (
	"fmt"
	"unsafe"
)

func main() {
	is := 1
	fmt.Println(&is)
	fmt.Println(*(&is))

	//有符号整数
	var intA int = -18 //与操作系统位数保持一致
	var intB int8 = 12

	fmt.Println("int")
	fmt.Printf("变量intA， 类型：%T, 字节：%d \n", intA, unsafe.Sizeof(intA))
	fmt.Printf("变量intB， 类型：%T, 字节：%d \n", intB, unsafe.Sizeof(intB))

	//无符号整数
	var uintA uint = 12 //与操作系统位数保持一致
	var uintB uint8 = 12

	fmt.Println("unit")
	fmt.Printf("变量uintA， 类型：%T, 字节：%d \n", uintA, unsafe.Sizeof(uintA))
	fmt.Printf("变量uintB， 类型：%T, 字节：%d \n", uintB, unsafe.Sizeof(uintB))

	fmt.Println("float")
	var floatA float32 = 12.2
	var floatB float64 = 13.3 //默认 float64
	fmt.Printf("变量floatA， 类型：%T, 字节：%d \n", floatA, unsafe.Sizeof(floatA))
	fmt.Printf("变量floatB， 类型：%T, 字节：%d \n", floatB, unsafe.Sizeof(floatB))

	fmt.Println("complex")
	c1 := complex(5, 7)
	c2 := 8 + 27i

	fmt.Printf("变量c1， 类型：%T, 字节：%d \n", c1, unsafe.Sizeof(c1))
	fmt.Printf("变量c2， 类型：%T, 字节：%d \n", c2, unsafe.Sizeof(c2))

	fmt.Println("byte 是unit8的别名，rune是int32的别名")
	//byte is an alias of uint8
	//rune is an alias of int32
	var byteA byte = 12
	var runeA rune = 23
	fmt.Printf("变量byteA， 类型：%T, 字节：%d \n", byteA, unsafe.Sizeof(byteA))
	fmt.Printf("变量runeA， 类型：%T, 字节：%d \n", runeA, unsafe.Sizeof(runeA))

	fmt.Println("string")
	first := "JIa"
	last := "Mengwei"
	name := first + last
	fmt.Printf("My name is %s \n", name)

	fmt.Println("类型转换")
	var a int = 1
	var b float64 = 23.9
	//c := a + b	//a 和 b 类型不一致无法运算
	c := a + int(b) //转换b的类型为int
	fmt.Println("c is ", c)

	var d = 1
	//var e float64 = d	//d 和 e类型不同无法赋值
	var e float64 = float64(d)
	fmt.Println("e is ", e)

}
