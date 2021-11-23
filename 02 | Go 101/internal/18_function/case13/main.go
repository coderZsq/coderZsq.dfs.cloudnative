package main

import "fmt"

func main() {
	// 此函数返回一个函数类型的结果，亦即闭包（closure）。
	isMultipleOfX := func (x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}

	var isMultipleOf3 = isMultipleOfX(3)
	var isMultipleOf5 = isMultipleOfX(5)
	fmt.Println(isMultipleOf3(6))  // true
	fmt.Println(isMultipleOf3(8))  // false
	fmt.Println(isMultipleOf5(10)) // true
	fmt.Println(isMultipleOf5(12)) // false

	isMultipleOf15 := func(n int) bool {
		return isMultipleOf3(n) && isMultipleOf5(n)
	}
	fmt.Println(isMultipleOf15(32)) // false
	fmt.Println(isMultipleOf15(60)) // true
}
