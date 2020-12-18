package main

import "fmt"

/**
method是一种特殊的func，它可以和一种特定的类型绑定

method receiver的指针参数可以接口指针和对象，func不可以
 */
func main() {
	//实例化一个employee
	emp := employee{name: "Jack"}
	//调用employee的about方法
	emp.about()
	about(emp)

	boss := boss{name: "Paul"}
	fmt.Println(boss)
	boss.changeName1("John1")
	fmt.Println(boss)
	//(&boss).changeName2("John2")  等价于下面这句， 因为编译重写boss为&boss
	boss.changeName2("John2")
	fmt.Println(boss)

	adrs := address{
		address: "Hangzhou",
	}

	user := user{
		name: "Lucuy",
		address : adrs,
	}

	user.about()
	user.printAddress()//直接访问类型字段的方法
	user.address.printAddress()//也可以使用完整写法
}

type employee struct {
	name string
}

//为employee类型声明一个方法about
func (e employee) about() {
	fmt.Println("my salary is ", e.name)
}

//用普通func也可以实现相同的功能
func about(e employee){
	fmt.Println("my salary is ", e.name)
}

//func 与 method的区别
//method可以与一个类型关联，我们可以将一组与特定类型相关的方法与类型进行关联，更好的面对对象
//相同的method名称可以定义与不同的类型进行关联

type boss struct{
	name string
}

func (b boss) about(){
	fmt.Println("I am a boss, my name is ", b.name)
}


//Pointer Receivers vs Value Receivers
func (b boss) changeName1(name string) {
	//不会改变boss的name
	b.name = name
}

func (b *boss) changeName2(name string) {
	//会改变boss的name
	b.name = name
}

//When to use pointer receiver and when to use value receiver
//1. 需要修改类型值得时候用指针，
//2. 类型实例占用空间较大，需要减小赋值类型实例时用指针
//3. 其他情况不用使用指针


type address struct {
	address string
}
func (adrs address) printAddress(){
	fmt.Println("address is ", adrs.address)
}

type  user struct {
	name string
	address
}

func (u user) about(){
	fmt.Println("I am ", u.name)
}

//无法在基本类型上定义method
//func (i int) haha(){
//
//}

type myint int

func (myi myint) haha(){
	fmt.Println("I am myint:", myi)
}

