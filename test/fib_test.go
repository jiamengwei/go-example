package main

import (
	"testing"
)

//用于代码检查测试
var name string

func TestFib(t *testing.T) {

	fibMap := make(map[int]int)
	fibMap[0] = 0
	fibMap[1] = 1
	fibMap[2] = 1
	fibMap[3] = 2
	fibMap[4] = 3
	fibMap[5] = 5
	fibMap[6] = 8
	fibMap[7] = 13
	fibMap[8] = 21
	fibMap[9] = 34

	for i, v := range fibMap {
		if r := Fib(i); r != v {
			t.Errorf("计算错误，期待结果：%d, 实际结果%d", v, r)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}
