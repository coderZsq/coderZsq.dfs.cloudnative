package main

func main() {
	var _ *int  // 一个基类型为int的非定义指针类型。
	var _ **int // 一个多级非定义指针类型，它的基类型为*int。

	type Ptr *int // Ptr是一个定义的指针类型，它的基类型为int。
	type PP *Ptr  // PP是一个定义的多级指针类型，它的基类型为Ptr。
}