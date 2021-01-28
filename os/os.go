package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	str, _ := os.Getwd()
	fmt.Println(str)
	fmt.Printf("%0.1s", "123456")

	sli := make([]string, 2)
	sli[0] = "士大夫"
	sli[1] = "s的风格"
	fmt.Println(sli)
	marshal, _ := json.Marshal(sli)
	fmt.Println(string(marshal))

	fmt.Println(strings.Join(sli, ""))
}
