package main

func main() {
	// 第一种方法（保持剩余元素的次序）：
	s = append(s[:i], s[i+1:]...)

	// 第二种方法（保持剩余元素的次序）：
	s = s[:i + copy(s[i:], s[i+1:])]

	// 上面两种方法都需要复制len(s)-i-1个元素。

	// 第三种方法（不保持剩余元素的次序）：
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]

	s[len(s):len(s)+1][0] = t0
	// 或者
	s[:len(s)+1][len(s)] = t0
}
