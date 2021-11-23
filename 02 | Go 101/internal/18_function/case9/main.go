package main

// 这两个函数的声明见前面几例。
func Sum(values ...int64) (sum int64) {......}
func Concat(sep string, tokens ...string) string {......}

func main() {
	// 下面两行报同样的错：实参数目太多了。
	_ = Sum(2, []int64{3, 5}...)
	_ = Concat(",", "Go", []string{"C", "Rust"}...)
}

