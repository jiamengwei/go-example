package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//file, err := ioutil.ReadFile("C:\\Users\\JiaMengwei\\Desktop")
	openFile, _ := os.OpenFile("C:\\Users\\JiaMengwei\\Desktop\\base64.txt", os.O_RDONLY, 0666)

	all, _ := ioutil.ReadAll(openFile)
	fmt.Println(string(all))
}
