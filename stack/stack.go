package main

import (
	"container/list"
	"errors"
	"fmt"
)

type Stack struct {
	list *list.List
}

func (stk *Stack) peek() (int, error) {
	val := stk.list.Front().Value
	i, ok := val.(int)
	if ok {
		return i, nil
	}
	return 0, errors.New("not found")
}

func (stk *Stack) push(val int) (int, error) {
	stk.list.PushFront(val)
	return val, nil
}

func (stk *Stack) pull() (int, error) {
	e := stk.list.Front()
	if e == nil {
		return 0, errors.New("not found")
	}
	stk.list.Remove(e)
	return e.Value.(int), nil
}

func (stk *Stack) toString() {
	l := stk.list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	stack := Stack{
		list: list.New(),
	}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.toString()
	fmt.Println(stack.pull())
	stack.toString()
}
