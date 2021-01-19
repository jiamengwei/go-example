package main

import (
	"fmt"
	"regexp"
)

func main() {
	regexpByServiceName, _ := regexp.Compile("service_name.*")
	match := regexpByServiceName.MatchString("service_name: user")
	fmt.Println(match)

	regexp, _ := regexp.Compile("=")
	arrays := regexp.Split("1A=2B=3C=4D=5E=6G=7Z", -1)
	for i, v := range arrays {
		fmt.Println(i, ":", v)
	}
}
