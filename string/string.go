package main

import (
	"fmt"
	"unicode/utf8"
)

/**
在UTF-8编码中，一个码位可以占用超过1个字节。
按字节打印一个字符时，会出现乱码现象
rune表示一个Unicode码位，不管一个码位包含多少字节都可以完整表示
*/
func main() {
	var firstName string
	firstName = "paul"
	fmt.Println(firstName)

	lastName := "Haw"
	fmt.Println(lastName)

	strA := "abcd"
	for i := 0; i < len(strA); i++ {
		fmt.Printf("%x , %c \n", strA[i], strA[i])
	}
	//61 , a
	//62 , b
	//63 , c
	//64 , d

	//在UTF-8编码中，一个码位可以占用超过1个字节。
	strB := "Señor"
	for i := 0; i < len(strB); i++ {
		fmt.Printf("%x , %c \n", strB[i], strB[i])
	}
	//53 , S
	//65 , e
	//c3 , Ã
	//b1 , ±
	//6f , o
	//72 , r
	runes := []rune("加上了雕刻技法")
	for i, v := range runes {
		fmt.Printf("i: %d, v:%c \n", i, v)
	}
	//i: 0, v:S
	//i: 1, v:e
	//i: 2, v:ñ
	//i: 3, v:o
	//i: 4, v:r

	byteSliceA := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	strC := string(byteSliceA)
	fmt.Println("strC is", strC)

	byteSliceB := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	strD := string(byteSliceB)
	fmt.Println("strD is", strD)
	fmt.Println("strC == strD ? ", strC == strD)
	fmt.Printf("strD number of bytes:%d, len:%d \n", len(strD), utf8.RuneCountInString(strD))

	strE := "hello"
	//strE[0] = 'a'  string不可变
	runeA := []rune(strE)
	runeA[0] = 'a'
	strE = string(runeA)
	fmt.Println("strE is", strE)

	str := "world"
	changeStr(str)
	fmt.Println(str)

	changeStrByPoint(&str)
	fmt.Println(str)
}

func changeStr(str string) {
	str = "hello"
}

func changeStrByPoint(str *string) {
	*str = "hello"
}
