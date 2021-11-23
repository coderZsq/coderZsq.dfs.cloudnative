package main

// 假设元素类型为T。
func Copy(dest, src []T) int {
	if len(dest) < len(src) {
		_ = append(dest[:0], src[:len(dest)]...)
		return len(dest)
	} else {
		_ = append(dest[:0], src...)
		return len(src)
	}
}
