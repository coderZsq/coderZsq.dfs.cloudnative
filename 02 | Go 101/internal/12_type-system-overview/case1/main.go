package main

func main() {
	// 假设T为任意一个类型，Tkey为一个支持比较的类型。
	type T string
	type Tkey int

	var _ *T         // 一个指针类型
	var _ [5]T       // 一个元素类型为T、元素个数为5的数组类型
	var _ []T        // 一个元素类型为T的切片类型
	var _ map[Tkey]T // 一个键值类型为Tkey、元素类型为T的映射类型

	// 一个结构体类型
	var _ struct {
		name string
		age  int
	}

	// 一个函数类型
	var _ func(int) (bool, string)

	// 一个接口类型
	var _ interface {
	    Method0(string) int
		Method1() (int, bool)
	}

	// 几个通道类型
	var _ chan T
	var _ chan<- T
	var _ <-chan T
}