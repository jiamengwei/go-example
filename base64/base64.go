package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("start----------------")
	//fileToBase64()
	base64ToFile()
	fmt.Println("end---------------")
}

func base64ToFile() {
	data, err := ioutil.ReadFile("D://base64.txt")
	if err != nil {
		panic(err)
	}
	ciphertext := strings.Replace(string(data), " ", "\n", -1)
	decodeData, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile("D://base64.jpg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(decodeData)
}
