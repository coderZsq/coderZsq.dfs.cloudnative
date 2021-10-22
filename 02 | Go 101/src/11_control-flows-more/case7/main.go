package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func(i int) {
				// 此i为形参i，非实参循环变量i。
				fmt.Println("b:", i)
			}(i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			i := i // 在下面的调用中，左i遮挡了右i。
			// <=> var i = i
			defer func() {
				// 此i为上面的左i，非循环变量i。
				fmt.Println("b:", i)
			}()
		}
	}()
}
