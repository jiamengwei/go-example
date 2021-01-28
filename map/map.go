package main

import "fmt"

func main() {
	//使用make声明一个map
	userAge := make(map[string]int)
	userAge["Jack"] = 12
	userAge["Lucy"] = 13
	fmt.Println(userAge)

	//声明时直接赋值
	userSex := map[string]string{
		"steve": "man",
		"jamie": "woman",
	}
	userSex["john"] = "man"
	fmt.Println(userSex)

	//使用下面的语句进行map的声明不会分配内存无法直接使用
	var userMapA map[string]int
	//userMapA["paul"] = 23	// panic: assignment to entry in nil map
	fmt.Println(userMapA, userMapA == nil)

	name := "steve"
	sex := userSex[name]
	fmt.Println("name:", name, "sex:", sex)

	sex = userSex["hahaha"]
	fmt.Println("hahaha's sex is", sex)

	age := userAge["hahaha"]
	fmt.Println("hahaha's age is", age)

	//判断一个元素是否存在
	name = "Jack"
	value, ok := userAge[name]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println(name, "not found")
	}

	//map的遍历
	for key, value := range userAge {
		fmt.Printf("key: %s, value: %d \n", key, value)
	}
	//删除指定元素
	delete(userAge, "Jack")
	delete(userAge, "sfasdfsdfasdf")
	fmt.Println(userAge)

	//使用map保存结构体
	type user struct {
		name string
		age  int
	}

	userA := user{
		name: "Jia",
		age:  14,
	}

	userB := user{
		name: "Wang",
		age:  45,
	}

	userMapB := map[string]user{
		"china": userA,
		"Japan": userB,
	}

	for country, userInfo := range userMapB {
		fmt.Printf("country :%s, name:%s, age:%d \n", country, userInfo.name, userInfo.age)
	}

	fmt.Println("userMapB length is", len(userMapB))

	//map是引用类型
	userMapC := userMapB
	userMapC["china"] = user{name: "Li", age: 23}
	fmt.Println(userMapB)

	//判断map是否相等
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

//map比较
func mapEquals(mapA, mapB map[string]int) bool {
	if len(mapA) != len(mapB) {
		return false
	}

	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}
	return true
}
