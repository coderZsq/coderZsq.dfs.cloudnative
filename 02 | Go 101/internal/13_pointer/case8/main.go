package main

func main() {
	type MyInt int64
	type Ta    *int64
	type Tb    *MyInt

	// 4个不同类型的指针：
	var pa0 Ta
	var pa1 *int64
	var pb0 Tb
	var pb1 *MyInt

	// 下面这6行编译没问题。它们的比较结果都为true。
	_ = pa0 == pa1
	_ = pb0 == pb1
	_ = pa0 == nil
	_ = pa1 == nil
	_ = pb0 == nil
	_ = pb1 == nil

	// 下面这三行编译不通过。
	/*
		_ = pa0 == pb0
		_ = pa1 == pb1
		_ = pa0 == Tb(nil)
	*/
}