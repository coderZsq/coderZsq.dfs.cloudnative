package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}
	// 改为一个切片。
	persons := []Person {{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 这次，此修改将反映在此次遍历过程中。
		persons[1].name = "Jack"
		// 这个修改仍然不会体现在persons切片容器中。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
}
