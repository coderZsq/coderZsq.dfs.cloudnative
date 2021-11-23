package main

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

func main() {
	
}
