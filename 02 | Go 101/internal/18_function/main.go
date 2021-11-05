package main

import "fmt"

func _() {
	//func (a int, b, c string) (x, y int, z bool)
	//func (x int, y, z string) (a, b int, c bool)
	//func (_ int, _, _ string) (_, _ int, _ bool)
	//
	//func (int, string, string) (int, int, bool) // 标准函数字面形式
	//func (a int, b string, c string) (int, int, bool)
	//func (x int, _ string, z string) (int, int, bool)
	//func (int, string, string) (x int, y int, z bool)
	//func (int, string, string) (a int, b int, _ bool)
}

func _() {
	//// 这三个函数类型字面形式是等价的。
	//func () (x int)
	//func () (int)
	//func () int
	//
	//// 这两个函数类型字面形式是等价的。
	//func (a int, b string) ()
	//func (a int, b string)
}

func _() {
	//func (values ...int64) (sum int64)
	//func (sep string, tokens ...string) string
}

func _() {
	//func Double(n int) (result int)
}

// Sum返回所有输入实参的和。
func Sum(values ...int64) (sum int64) {
	// values的类型为[]int64。
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}

// Concat是一个低效的字符串拼接函数。
func Concat(sep string, tokens ...string) string {
	// tokens的类型为[]string。
	r := ""
	for i, t := range tokens {
		if i != 0 {
			r += sep
		}
		r += t
	}
	return r
}

func _() {
	//func Print(a ...interface{}) (n int, err error)
	//func Printf(format string, a ...interface{}) (n int, err error)
	//func Println(a ...interface{}) (n int, err error)
}

func init() {
	a0 := Sum()
	a1 := Sum(2)
	a3 := Sum(2, 3, 5)
	// 上面三行和下面三行是等价的。
	b0 := Sum([]int64{}...) // <=> Sum(nil...)
	b1 := Sum([]int64{2}...)
	b3 := Sum([]int64{2, 3, 5}...)
	fmt.Println(a0, a1, a3) // 0 2 10
	fmt.Println(b0, b1, b3) // 0 2 10
}

func init() {
	tokens := []string{"Go", "C", "Rust"}
	langsA := Concat(",", tokens...)         // 风格1
	langsB := Concat(",", "Go", "C", "Rust") // 风格2
	fmt.Println(langsA == langsB)            // true
}

func fa() int {
a:
	goto a
}

func fb() bool {
	for {
	}
}

func HalfAndNegative(n int) (int, int) {
	return n / 2, -n
}

func AddSub(a, b int) (int, int) {
	return a + b, a - b
}

func Dummy(values ...int) {}

func init() {
	// 这几行编译没问题。
	AddSub(HalfAndNegative(6))
	AddSub(AddSub(AddSub(7, 5)))
	AddSub(AddSub(HalfAndNegative(6)))
	Dummy(HalfAndNegative(6))
	_, _ = AddSub(7, 5)

	// 下面这几行编译不通过。
	/*
		_, _, _ = 6, AddSub(7, 5)
		Dummy(AddSub(7, 5), 9)
		Dummy(AddSub(7, 5), HalfAndNegative(6))
	*/
}

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n) // f的类型为"func(int) int"
}

func init() {
	fmt.Printf("%T\n", Double) // func(int) int
	// Double = nil // error: Double是不可修改的

	var f func(n int) int // 默认值为nil
	f = Double
	g := Apply
	fmt.Printf("%T\n", g) // func(int, func(int) int) int

	fmt.Println(f(9))         // 18
	fmt.Println(g(6, Double)) // 12
	fmt.Println(Apply(6, f))  // 12
}

func init() {
	// 此函数返回一个函数类型的结果，亦即闭包（closure）。
	isMultipleOfX := func(x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}

	var isMultipleOf3 = isMultipleOfX(3)
	var isMultipleOf5 = isMultipleOfX(5)
	fmt.Println(isMultipleOf3(6))  // true
	fmt.Println(isMultipleOf3(8))  // false
	fmt.Println(isMultipleOf5(10)) // true
	fmt.Println(isMultipleOf5(12)) // false

	isMultipleOf15 := func(n int) bool {
		return isMultipleOf3(n) && isMultipleOf5(n)
	}
	fmt.Println(isMultipleOf15(32)) // false
	fmt.Println(isMultipleOf15(60)) // true
}

func main() {

}
