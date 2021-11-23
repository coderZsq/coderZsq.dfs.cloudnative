package main

func main() {
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
