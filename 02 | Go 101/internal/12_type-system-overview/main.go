package main

func init() {
	//// 假设T为任意一个类型，Tkey为一个支持比较的类型。
	//
	//*T         // 一个指针类型
	//[5]T       // 一个元素类型为T、元素个数为5的数组类型
	//[]T        // 一个元素类型为T的切片类型
	//map[Tkey]T // 一个键值类型为Tkey、元素类型为T的映射类型
	//
	//// 一个结构体类型
	//struct {
	//	name string
	//	age  int
	//}
	//
	//// 一个函数类型
	//func(int) (bool, string)
	//
	//// 一个接口类型
	//interface {
	//	Method0(string) int
	//	Method1() (int, bool)
	//}
	//
	//// 几个通道类型
	//chan T
	//chan<- T
	//<-chan T
}

func init() {
	//// 定义单个类型。
	//type NewTypeName SourceType
	//
	//// 定义多个类型。
	//type (
	//	NewTypeName1 SourceType1
	//	NewTypeName2 SourceType2
	//)
}

func init() {
	// 下面这些新定义的类型和它们的源类型都是基本类型。
	type (
		MyInt int
		Age   int
		Text  string
	)

	// 下面这些新定义的类型和它们的源类型都是组合类型。
	type IntPtr *int
	type Book struct{author, title string; pages int}
	type Convert func(in0 int, in1 bool)(out0 int, out1 string)
	type StringArray [5]string
	type StringSlice []string
}

func f() {
	// 这三个新定义的类型名称只能在此函数内使用。
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{Read([]byte) int}
}

func init() {
	type (
		Name = string
		Age  = int
	)

	type table = map[string]int
	type Table = map[Name]Age
}

func init() {
	type A []string
	type B = A
	type C = []string
}

func init() {
	// 这四个类型的底层类型均为内置类型int。
	type (
		MyInt int
		Age   MyInt
	)

	// 下面这三个新声明的类型的底层类型各不相同。
	type (
		IntSlice   []int   // 底层类型为[]int
		MyIntSlice []MyInt // 底层类型为[]MyInt
		AgeSlice   []Age   // 底层类型为[]Age
	)

	// 类型[]Age、Ages和AgeSlice的底层类型均为[]Age。
	type Ages AgeSlice
}

func init() {
	//struct {
	//	author string
	//	title  string
	//	pages  int
	//}
}