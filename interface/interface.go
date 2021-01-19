package main

import "fmt"

/**
接口是一组方法签名的集合，当一个类型定义了一个接口的所有方法签名时，我们就可以说这个类型实现了该接口
接口至此指针实现和值实现
接口支持嵌套
*/
type a interface {
	funcA() int
}

type b interface {
	funcB() int
}

type ab interface {
	a
	b
}

//声明一个接口SalaryCalculator，在接口中定义一个方法
type SalaryCalculator interface {
	CalculateSalary() int
}

//声明一个employee类型
type employee struct {
	baseSalary int
	bonus      int
}

//为employee定义接口SalaryCalculator中的方法，此时的employee类型即实现了SalaryCalculator接口
func (epl employee) CalculateSalary() int {
	return epl.baseSalary + epl.bonus
}

type seniorEmployee struct {
	baseSalary int
	bonus      int
	commission int
}

func (sepl seniorEmployee) CalculateSalary() int {
	return sepl.baseSalary + sepl.bonus + sepl.commission
}

//没有定义方法签名的接口叫做方法签名，因此所有的类型都实现了空接口
func describe(inter interface{}) {
	//打印接口的具体类型
	fmt.Printf("type of inter: %T \n", inter)
}

//类型断言用于提取接口的底层值
//i.(T)是用于获取接口i的底层值的语法，接口i的具体类型是T
func assert(i interface{}) {
	//v, ok := i.(int)

	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	case employee:
		fmt.Printf("I am an employee and my value is %d\n", i.(employee))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	eplA := employee{
		baseSalary: 1000,
		bonus:      10,
	}

	salaryA := eplA.CalculateSalary()
	fmt.Println(salaryA)

	salaryCalculator := seniorEmployee{
		baseSalary: 1200,
		bonus:      20,
		commission: 500,
	}
	fmt.Println(salaryCalculator.CalculateSalary())

	describe(1)
	describe("A")
	describe([]int{1, 2, 3})
	//type of inter: int
	//type of inter: string
	//type of inter: []int
	assert(1)
	assert(12)
	assert("A")
	assert(eplA)

}
