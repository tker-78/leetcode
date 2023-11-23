package main

import "fmt"

type MyStack struct {
	values []int
}

func Constructor() MyStack {
	ms := MyStack{}
	return ms
}

func (ms *MyStack) Push(x int) {
	ms.values = append(ms.values, x)
}

func (ms *MyStack) Pop() int {
	lenValues := len(ms.values)
	if lenValues == 0 {
		return 0
	}

	if lenValues == 1 {
		val := ms.values[lenValues-1]
		ms.values = []int{}
		return val
	}

	val := ms.values[lenValues-1]
	ms.values = ms.values[:lenValues-2]
	return val

}

func (ms *MyStack) Top() int {
	lenValues := len(ms.values)
	if lenValues == 0 {
		return 0
	}
	return ms.values[lenValues-1]
}

func (ms *MyStack) Empty() bool {
	return len(ms.values) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	para1 := obj.Pop()
	para2 := obj.Top()
	para3 := obj.Empty()

	fmt.Println(para1, para2, para3)
}
