package employee

import "fmt"

/**
不直接暴露employee而是使用New方法来构造一个employee对象来避免使用employee时未经初始化
*/

type employee struct {
	name string
	age  int
}

func (e employee) Info() {
	fmt.Printf("name:%s, age:%d \n", e.name, e.age)
}

func New(name string, age int) employee {
	return employee{name: name, age: age}
}
