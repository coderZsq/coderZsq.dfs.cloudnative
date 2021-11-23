package main

type Encoder interface {Encode([]byte) []byte}
type Person struct {name string; age int}
type Alias = struct {name string; age int}
type AliasPtr = *struct {name string; age int}
type IntPtr *int
type AliasPP = *IntPtr

func main() {
	var _ struct {
		// 这些类型或别名都可以被内嵌。
		Encoder
		Person
		*Person
		Alias
		*Alias
		AliasPtr
		int
		*int
	}

	var _ struct{
		// 这些类型或别名都不能被内嵌。
		AliasPP          // 基类型为一个指针类型
		*Encoder         // 基类型为一个接口类型
		*AliasPtr        // 基类型为一个指针类型
		IntPtr           // 定义的指针类型
		*IntPtr          // 基类型为一个指针类型
		*chan int        // 基类型为一个非定义类型
		struct {age int} // 非定义非指针类型
		map[string]int   // 非定义非指针类型
		[]int64          // 非定义非指针类型
		func()           // 非定义非指针类型
	}
}