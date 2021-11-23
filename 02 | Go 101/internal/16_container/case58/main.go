package main

func main() {
	// 前弹出（pop front，又称shift）
	s, e = s[1:], s[0]
	// 后弹出（pop back）
	s, e = s[:len(s)-1], s[len(s)-1]
	// 前推（push front）
	s = append([]T{e}, s...)
	// 后推（push back）
	s = append(s, e)
}
