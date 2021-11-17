package main

import "fmt"

func init() {
	// 一个非定义接口类型。
	//interface {
	//	About() string
	//}

	//ReadWriter是一个定义的接口类型。
	type ReadWriter interface {
		Read(buf []byte) (n int, err error)
		Write(buf []byte) (n int, err error)
	}

	// Runnable是一个非定义接口类型的别名。
	type Runnable = interface {
		Run()
	}
}

func init() {
	type error interface {
		Error() string
	}
}

func init() {
	// 一个非定义空接口类型。
	//interface{}

	// 类型I是一个定义的空接口类型。
	type I interface{}
}

type Book struct {
	name string
	// 更多其它字段……
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "我是一个自定义整数类型"
}

func init() {
	//	import "database/sql"
	//...
	//
	//type DatabaseStorer interface {
	//	Exec(query string, args ...interface{}) (sql.Result, error)
	//	Prepare(query string) (*sql.Stmt, error)
	//	Query(query string, args ...interface{}) (*sql.Rows, error)
	//}
}

func init() {
	type Aboutable interface {
		About() string
	}

	// 一个*Book值被包裹在了一个Aboutable值中。
	var a Aboutable = &Book{"Go语言101"}
	fmt.Println(a) // &{Go语言101}

	// i是一个空接口值。类型*Book实现了任何空接口类型。
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	// Aboutable实现了空接口类型interface{}。
	i = a
	fmt.Println(i) // &{Go语言101}
}

func init() {
	// func Println(a ...interface{}) (n int, err error)
}

func init() {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i) // [1 2 3]
	i = map[string]int{"Go": 2012}
	fmt.Println(i) // map[Go:2012]
	i = true
	fmt.Println(i) // true
	i = 1
	fmt.Println(i) // 1
	i = "abc"
	fmt.Println(i) // abc

	// 将接口值i中包裹的值清除掉。
	i = nil
	fmt.Println(i) //
}

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter用来删除重复的数字。
type UniqueFilter struct{}
func (UniqueFilter) About() string {
	return "删除重复的数字"
}
func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

// MultipleFilter用来只保留某个整数的倍数数字。
type MultipleFilter int
func (mf MultipleFilter) About() string {
	return fmt.Sprintf("保留%v的倍数", mf)
}
func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n % int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

// 在多态特性的帮助下，只需要一个filteAndPrint函数。
func filteAndPrint(fltr Filter, unfiltered []int) []int {
	// 在fltr参数上调用方法其实是调用fltr的动态值的方法。
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About() + ":\n\t", filtered)
	return filtered
}

func init() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("过滤之前：\n\t", numbers)

	// 三个非接口值被包裹在一个Filter切片
	// 的三个接口元素中。
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	// 每个切片元素将被赋值给类型为Filter的
	// 循环变量fltr。每个元素中的动态值也将
	// 被同时复制并被包裹在循环变量fltr中。
	for _, fltr := range filters {
		numbers = filteAndPrint(fltr, numbers)
	}
}

func init() {
	// 编译器将把123的类型推断为内置类型int。
	var x interface{} = 123

	// 情形一：
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// 情形二：
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// 情形三：
	//a = x.(float64) // 将产生一个恐慌
}

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}
func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func init() {
	var x interface{} = DummyWriter{}
	// y的动态类型为内置类型string。
	var y interface{} = "abc"
	var w Writer
	var ok bool

	// DummyWriter既实现了Writer，也实现了interface{}。
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// y的动态类型为string。string类型并没有实现Writer。
	w, ok = y.(Writer)
	fmt.Println(w, ok) //  false
	//w = y.(Writer)     // 将产生一个恐慌
}

func init() {
	//switch aSimpleStatement; v := x.(type) {
	//case TypeA:
	//	...
	//case TypeB, TypeC:
	//	...
	//case nil:
	//	...
	//default:
	//	...
	//}
}

func init() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// 这里，虽然变量v只被声明了一次，但是它在
		// 不同分支中可以表现为多个类型的变量值。
		switch v := x.(type) {
		case []int: // 一个类型字面表示
			// 在此分支中，v的类型为[]int。
			fmt.Println("int slice:", v)
		case string: // 一个类型名
			// 在此分支中，v的类型为string。
			fmt.Println("string:", v)
		case int, float64, int32: // 多个类型名
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("number:", v)
		case nil:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println(v)
		default:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("others:", v)
		}
		// 注意：在最后的三个分支中，v均为接口值x的一个复制。
	}
}

func init() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		if v, ok := x.([]int); ok {
			fmt.Println("int slice:", v)
		} else if v, ok := x.(string); ok {
			fmt.Println("string:", v)
		} else if x == nil {
			v := x
			fmt.Println(v)
		} else {
			_, isInt := x.(int)
			_, isFloat64 := x.(float64)
			_, isInt32 := x.(int32)
			if isInt || isFloat64 || isInt32 {
				v := x
				fmt.Println("number:", v)
			} else {
				v := x
				fmt.Println("others:", v)
			}
		}
	}
}

func init() {
	type Ia interface {
		fa()
	}

	type Ib = interface {
		fb()
	}

	type Ic interface {
		fa()
		fb()
	}

	type Id = interface {
		Ia // 内嵌Ia
		Ib // 内嵌Ib
	}

	type Ie interface {
		Ia // 内嵌Ia
		fb()
	}

	type Ix interface {
		Ia
		Ic
	}

	type Iy = interface {
		Ib
		Ic
	}

	type Iz interface {
		Ic
		fa()
	}
}

func init() {
	var a, b, c interface{} = "abc", 123, "a"+"b"+"c"
	fmt.Println(a == b) // 第二步的情形。输出"false"。
	fmt.Println(a == c) // 第三步的情形。输出"true"。

	var x *int = nil
	var y *bool = nil
	var ix, iy interface{} = x, y
	var i interface{} = nil
	fmt.Println(ix == iy) // 第二步的情形。输出"false"。
	fmt.Println(ix == i)  // 第一步的情形。输出"false"。
	fmt.Println(iy == i)  // 第一步的情形。输出"false"。

	var s []int = nil // []int为一个不可比较类型。
	i = s
	fmt.Println(i == nil) // 第一步的情形。输出"false"。
	fmt.Println(i == i)   // 第三步的情形。将产生一个恐慌。
}

type I interface {
	m(int)bool
}

type T string
func (t T) m(n int) bool {
	return len(t) > n
}

func init() {
	var i I = T("gopher")
	fmt.Println(i.m(5))                          // true
	fmt.Println(I.m(i, 5))                       // true
	//fmt.Println(interface {m(int) bool}.m(i, 5)) // true

	// 下面这几行被执行的时候都将会产生一个恐慌。
	//I(nil).m(5)
	//I.m(nil, 5)
	//interface {m(int) bool}.m(nil, 5)
}

func main() {

}
