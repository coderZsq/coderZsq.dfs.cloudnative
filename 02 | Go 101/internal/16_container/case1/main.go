package main

func main() {
	const Size = 32

	type Person struct {
		name string
		age  int
	}

	// 数组类型
	var _ [5]string
	var _ [Size]int
	var _ [16][]byte  // 元素类型为一个切片类型：[]byte
	var _ [100]Person // 元素类型为一个结构体类型：Person

	// 切片类型
	var _ []bool
	var _ []int64
	var _ []map[int]bool // 元素类型为一个映射类型：map[int]bool
	var _ []*int         // 元素类型为一个指针类型：*int

	// 映射类型
	var _ map[string]int
	var _ map[int]bool
	var _ map[int16][6]string     // 元素类型为一个数组类型：[6]string
	var _ map[bool][]string       // 元素类型为一个切片类型：[]string
	var _ map[struct{x int}]*int8 // 元素类型为一个指针类型：*int8；
	// 键值类型为一个结构体类型。
}