package main

func main() {
	// 第一种方法:单行实现。
	s = append(s[:i], append(elements, s[i:]...)...)

	// 上面这种单行实现把s[i:]中的元素复制了两次，并且它可能
	// 最多导致两次内存开辟（最少一次）。
	// 下面这种繁琐的实现只把s[i:]中的元素复制了一次，并且
	// 它最多只会导致一次内存开辟（最少零次）。
	// 但是，在当前的官方标准编译器实现中（1.17版本），此
	// 繁琐实现中的make调用将会把一些刚开辟出来的元素清零。
	// 这其实是没有必要的。所以此繁琐实现并非总是比上面的
	// 单行实现效率更高。事实上，它仅在处理小切片时更高效。

	if cap(s) >= len(s) + len(elements) {
		s = s[:len(s)+len(elements)]
		copy(s[i+len(elements):], s[i:])
		copy(s[i:], elements)
	} else {
		x := make([]T, 0, len(elements)+len(s))
		x = append(x, s[:i]...)
		x = append(x, elements...)
		x = append(x, s[i:]...)
		s = x
	}

	// Push（插入到结尾）。
	s = append(s, elements...)

	// Unshift（插入到开头）。
	s = append(elements, s...)
}
