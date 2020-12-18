package main

import "fmt"

func main() {
	userAge := make(map[string]int)
	userAge["Jack"] = 12
	userAge["Lucy"] = 13
	fmt.Println(userAge)

	userSex := map[string]string{
		"steve" : "man",
		"jamie" : "woman",
	}
	userSex["john"] = "man"
	fmt.Println(userSex)

	var userMapA map[string]int
	//userMapA["paul"] = 23	// panic: assignment to entry in nil map
	fmt.Println(userMapA, userMapA == nil)

	name := "steve"
	sex := userSex[name]
	fmt.Println("name:", name, "sex:", sex)

	sex = userSex["hahaha"]
	fmt.Println("sex is",sex)

	age := userAge["hahaha"]
	fmt.Println("age is",age)

	name = "Jack"
	value, ok := userAge[name]
	if ok {
		fmt.Println(value)
	}else {
		fmt.Println(name, "not found")
	}


	for key, value := range userAge{
		fmt.Printf("key: %s, value: %d \n",key, value)
	}

	delete(userAge,  "Jack")
	delete(userAge,  "sfasdfsdfasdf")
	fmt.Println(userAge)

	type user struct {
		name string
		age int
	}

	userA := user{
		name: "Jia",
		age: 14,
	}

	userB := user{
		name: "Wang",
		age : 45,
	}

	userMapB := map[string]user{
		"china" :  userA,
		"Japan" : userB,
	}

	for country, userInfo := range userMapB{
		fmt.Printf("country :%s, name:%s, age:%d \n", country, userInfo.name, userInfo.age)
	}

	fmt.Println("userMapB length is", len(userMapB))

	userMapC :=  userMapB
	userMapC["china"] = user{name:"Li",age:23}
	fmt.Println(userMapB)

	userAgeCopy := userAge
	equals := mapEquals(userAgeCopy, userAge)
	if equals {
		fmt.Println("userAgeCopy equals userAge")
	}
	equals = mapEquals(userMapA, userAge)
	if !equals {
		fmt.Println("userMapA not equals userAge")
	}

}

func mapEquals(mapA,mapB map[string]int) bool {
	if len(mapA) != len(mapB) {
		return false
	}

	for k ,v := range mapA{
		if mapB[k] != v {
			return false
		}
	}
	return true
}
