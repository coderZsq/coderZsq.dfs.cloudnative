package main

func main() {
	// "len(s)+to-from"是删除操作之前切片s的长度。
	temp := s[len(s):len(s)+to-from]
	for i := range temp {
		temp[i] = t0 // t0是类型T的零值字面量
	}
}
