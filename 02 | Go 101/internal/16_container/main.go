package main

import (
	"fmt"
	"unsafe"
)

func init() {
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
	var _ []map[int]bool // 元素类型为一个映射类型：	map[int]bool
	var _ []*int         // 元素类型为一个指针类型：*int

	// 映射类型
	var _ map[string]int
	var _ map[int]bool
	var _ map[int16][6]string       // 元素类型为一个数组类型：[6]string
	var _ map[bool][]string         // 元素类型为一个切片类型：[]string
	var _ map[struct{ x int }]*int8 // 元素类型为一个指针类型：*int8；
	// 键值类型为一个结构体类型。
}

func init() {
	// 一个含有4个布尔元素的数组值。
	_ = [4]bool{false, true, true, false}

	// 一个含有三个字符串值的切片值。
	_ = []string{"break", "continue", "fallthrough"}

	// 一个映射值。
	_ = map[string]int{"C": 1972, "Python": 1991, "Go": 2009}
}

func init() {
	// 下面这些切片字面量都是等价的。
	_ = []string{"break", "continue", "fallthrough"}
	_ = []string{0: "break", 1: "continue", 2: "fallthrough"}
	_ = []string{2: "fallthrough", 1: "continue", 0: "break"}
	_ = []string{2: "fallthrough", 0: "break", "continue"}

	// 下面这些数组字面量都是等价的。
	_ = [4]bool{false, true, true, false}
	_ = [4]bool{0: false, 1: true, 2: true, 3: false}
	_ = [4]bool{1: true, true}
	_ = [4]bool{2: true, 1: true}
	_ = [...]bool{false, true, true, false}
	_ = [...]bool{3: false, 1: true, true}
}

func init() {
	var a uint = 1
	_ = map[uint]int{a: 123} // 没问题
	//var _ = []int{a: 100}          // error: 下标必须为常量
	//var _ = [5]int{a: 100}         // error: 下标必须为常量
}

func init() {
	pm := &map[string]int{"C": 1972, "Go": 2009}
	ps := &[]string{"break", "continue"}
	pa := &[...]bool{false, true, true, false}
	fmt.Printf("%T\n", pm) // *map[string]int
	fmt.Printf("%T\n", ps) // *[]string
	fmt.Printf("%T\n", pa) // *[4]bool
}

func init() {
	// heads为一个切片值。它的类型的元素类型为*[4]byte。
	// 此元素类型为一个基类型为[4]byte的指针类型。
	// 此指针基类型为一个元素类型为byte的数组类型。
	var heads = []*[4]byte{
		&[4]byte{'P', 'N', 'G', ' '},
		&[4]byte{'G', 'I', 'F', ' '},
		&[4]byte{'J', 'P', 'E', 'G'},
	}
	_ = heads
}

func init() {
	var heads = []*[4]byte{
		{'P', 'N', 'G', ' '},
		{'G', 'I', 'F', ' '},
		{'J', 'P', 'E', 'G'},
	}
	_ = heads
}

func init() {
	type language struct {
		name string
		year int
	}

	_ = [...]language{
		language{"C", 1972},
		language{"Python", 1991},
		language{"Go", 2009},
	}
}

func init() {
	type language struct {
		name string
		year int
	}

	_ = [...]language{
		{"C", 1972},
		{"Python", 1991},
		{"Go", 2009},
	}
}

func init() {
	type LangCategory struct {
		dynamic bool
		strong  bool
	}

	// 此映射值的类型的键值类型为一个结构体类型，
	// 元素类型为另一个映射类型：map[string]int。
	var _ = map[LangCategory]map[string]int{
		LangCategory{true, true}: map[string]int{
			"Python": 1991,
			"Erlang": 1986,
		},
		LangCategory{true, false}: map[string]int{
			"JavaScript": 1995,
		},
		LangCategory{false, true}: map[string]int{
			"Go":   2009,
			"Rust": 2010,
		},
		LangCategory{false, false}: map[string]int{
			"C": 1972,
		},
	}
}

func init() {
	type LangCategory struct {
		dynamic bool
		strong  bool
	}

	var _ = map[LangCategory]map[string]int{
		{true, true}: {
			"Python": 1991,
			"Erlang": 1986,
		},
		{true, false}: {
			"JavaScript": 1995,
		},
		{false, true}: {
			"Go":   2009,
			"Rust": 2010,
		},
		{false, false}: {
			"C": 1972,
		},
	}
}

func init() {
	var a [16]byte
	var s []int
	var m map[string]int

	fmt.Println(a == a)   // true
	fmt.Println(m == nil) // true
	fmt.Println(s == nil) // true
	fmt.Println(nil == map[string]int{}) // false
	fmt.Println(nil == []int{})          // false

	// 下面这些行编译不通过。
	/*
		_ = m == m
		_ = s == s
		_ = m == map[string]int(nil)
		_ = s == []int(nil)
		var x [16][]int
		_ = x == x
		var y [16]map[int]bool
		_ = y == y
	*/
}

func init() {
	var a [5]int
	fmt.Println(len(a), cap(a)) // 5 5
	var s []int
	fmt.Println(len(s), cap(s)) // 0 0
	s, s2 := []int{2, 3, 5}, []bool{}
	fmt.Println(len(s), cap(s), len(s2), cap(s2)) // 3 3 0 0
	var m map[int]bool
	fmt.Println(len(m)) // 0
	m, m2 := map[int]bool{1: true, 0: false}, map[int]int{}
	fmt.Println(len(m), len(m2)) // 2 0
}

func init() {
	a := [3]int{-1, 0, 1}
	s := []bool{true, false}
	m := map[string]int{"abc": 123, "xyz": 789}
	fmt.Println (a[2], s[1], m["abc"])    // 读取
	a[2], s[1], m["abc"] = 999, true, 567 // 修改
	fmt.Println (a[2], s[1], m["abc"])    // 读取

	n, present := m["hello"]
	fmt.Println(n, present, m["hello"]) // 0 false 0
	n, present = m["abc"]
	fmt.Println(n, present, m["abc"]) // 567 true 567
	m = nil
	fmt.Println(m["abc"]) // 0

	// 下面这两行编译不同过。
	/*
		_ = a[3]  // 下标越界
		_ = s[-1] // 下标越界
	*/

	// 下面这几行每行都会造成一个恐慌。
	//_ = a[n]         // panic: 下标越界
	//_ = s[n]         // panic: 下标越界
	//m["hello"] = 555 // panic: m为一个零值映射
}

func init() {
	type _slice struct {
		elements unsafe.Pointer // 引用着底层存储在间接部分上的元素
		len      int            // 长度
		cap      int            // 容量
	}
}

func init() {
	m0 := map[int]int{0:7, 1:8, 2:9}
	m1 := m0
	m1[0] = 2
	fmt.Println(m0, m1) // map[0:2 1:8 2:9] map[0:2 1:8 2:9]

	s0 := []int{7, 8, 9}
	s1 := s0
	s1[0] = 2
	fmt.Println(s0, s1) // [2 8 9] [2 8 9]

	a0 := [...]int{7, 8, 9}
	a1 := a0
	a1[0] = 2
	fmt.Println(a0, a1) // [7 8 9] [2 8 9]
}

func init() {
	m := map[string]int{"Go": 2007}
	m["C"] = 1972     // 添加
	m["Java"] = 1995  // 添加
	fmt.Println(m)    // map[C:1972 Go:2007 Java:1995]
	m["Go"] = 2009    // 修改
	delete(m, "Java") // 删除
	fmt.Println(m)    // map[C:1972 Go:2009]
}

func init() {
	s0 := []int{2, 3, 5}
	fmt.Println(s0, cap(s0)) // [2 3 5] 3
	s1 := append(s0, 7)      // 添加一个元素
	fmt.Println(s1, cap(s1)) // [2 3 5 7] 6
	s2 := append(s1, 11, 13) // 添加两个元素
	fmt.Println(s2, cap(s2)) // [2 3 5 7 11 13] 6
	s3 := append(s0)         // <=> s3 := s0
	fmt.Println(s3, cap(s3)) // [2 3 5] 3
	s4 := append(s0, s0...)  // 以s0为基础添加s0中所有的元素
	fmt.Println(s4, cap(s4)) // [2 3 5 2 3 5] 6

	s0[0], s1[0] = 99, 789
	fmt.Println(s2[0], s3[0], s4[0]) // 789 99 2
}

func init() {
	var s = append([]string(nil), "array", "slice")
	fmt.Println(s)      // [array slice]
	fmt.Println(cap(s)) // 2
	s = append(s, "map")
	fmt.Println(s)      // [array slice map]
	fmt.Println(cap(s)) // 4
	s = append(s, "channel")
	fmt.Println(s)      // [array slice map channel]
	fmt.Println(cap(s)) // 4
}

func init() {
	// 创建映射。
	fmt.Println(make(map[string]int)) // map[]
	m := make(map[string]int, 3)
	fmt.Println(m, len(m)) // map[] 0
	m["C"] = 1972
	m["Go"] = 2009
	fmt.Println(m, len(m)) // map[C:1972 Go:2009] 2

	// 创建切片。
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5
	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s)) // [0 0] 2 2
}

func init() {
	m := *new(map[string]int)   // <=> var m map[string]int
	fmt.Println(m == nil)       // true
	s := *new([]int)            // <=> var s []int
	fmt.Println(s == nil)       // true
	a := *new([5]bool)          // <=> var a [5]bool
	fmt.Println(a == [5]bool{}) // true
}

func init() {
	a := [5]int{2, 3, 5, 7}
	s := make([]bool, 2)
	pa2, ps1 := &a[2], &s[1]
	fmt.Println(*pa2, *ps1) // 5 false
	a[2], s[1] = 99, true
	fmt.Println(*pa2, *ps1) // 99 true
	ps0 := &[]string{"Go", "C"}[0]
	fmt.Println(*ps0) // Go

	m := map[int]bool{1: true}
	_ = m
	// 下面这几行编译不通过。
	/*
		_ = &[3]int{2, 3, 5}[0]
		_ = &map[int]bool{1: true}[1]
		_ = &m[1]
	*/
}

func init() {
	type T struct{age int}
	mt := map[string]T{}
	mt["John"] = T{age: 29} // 整体修改是允许的
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789} // 整体修改是允许的

	// 这两个赋值编译不通过，因为部分修改一个映射
	// 元素是非法的。这看上去确实有些反直觉。
	/*
		ma[1][1] = 123      // error
		mt["John"].age = 30 // error
	*/

	// 读取映射元素的元素或者字段是没问题的。
	fmt.Println(ma[1][1])       // 789
	fmt.Println(mt["John"].age) // 29
}

func init() {
	type T struct{age int}
	mt := map[string]T{}
	mt["John"] = T{age: 29}
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789}

	t := mt["John"] // 临时变量
	t.age = 30
	mt["John"] = t // 整体修改

	a := ma[1] // 临时变量
	a[1] = 123
	ma[1] = a // 整体修改

	fmt.Println(ma[1][1], mt["John"].age) // 123 30
}

func init() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	s0 := a[:]     // <=> s0 := a[0:7:7]
	s1 := s0[:]    // <=> s1 := s0
	s2 := s1[1:3]  // <=> s2 := a[1:3]
	s3 := s1[3:]   // <=> s3 := s1[3:7]
	s4 := s0[3:5]  // <=> s4 := s0[3:5:7]
	s5 := s4[:2:2] // <=> s5 := s0[3:5:5]
	s6 := append(s4, 77)
	s7 := append(s5, 88)
	s8 := append(s7, 66)
	s3[1] = 99
	fmt.Println(len(s2), cap(s2), s2) // 2 6 [1 2]
	fmt.Println(len(s3), cap(s3), s3) // 4 4 [3 99 77 6]
	fmt.Println(len(s4), cap(s4), s4) // 2 4 [3 99]
	fmt.Println(len(s5), cap(s5), s5) // 2 2 [3 99]
	fmt.Println(len(s6), cap(s6), s6) // 3 4 [3 99 77]
	fmt.Println(len(s7), cap(s7), s7) // 3 4 [3 4 88]
	fmt.Println(len(s8), cap(s8), s8) // 4 4 [3 4 88 66]
}

func f() []int {
	s := make([]int, 10, 100)
	return s[50:60]
}

func init() {
	type S []int
	type A [2]int
	type P *A

	var x []int
	var y = make([]int, 0)
	var x0 = (*[0]int)(x) // okay, x0 == nil
	var y0 = (*[0]int)(y) // okay, y0 != nil
	_, _ = x0, y0

	var z = make([]int, 3, 5)
	var _ = (*[3]int)(z) // okay
	var _ = (*[2]int)(z) // okay
	var _ = (*A)(z)      // okay
	var _ = P(z)         // okay

	var w = S(z)
	var _ = (*[3]int)(w) // okay
	var _ = (*[2]int)(w) // okay
	var _ = (*A)(w)      // okay
	var _ = P(w)         // okay

	var _ = (*[4]int)(z) // 会产生恐慌
}

func main() {

}
