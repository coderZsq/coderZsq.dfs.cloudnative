package main

func main() {
	// 第一种方法（保持剩余元素的次序）：
	s = append(s[:from], s[to:]...)

	// 第二种方法（保持剩余元素的次序）：
	s = s[:from + copy(s[from:], s[to:])]

	// 第三种方法（不保持剩余元素的次序）：
	if n := to-from; len(s)-to < n {
		copy(s[from:to], s[to:])
	} else {
		copy(s[from:to], s[len(s)-n:])
	}
	s = s[:len(s)-(to-from)]
}
