package main

// 假设T是一个小尺寸类型。
func DeleteElements(s []T, keep func(T) bool, clear bool) []T {
	// result := make([]T, 0, len(s))
	result := s[:0] // 无须开辟内存
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	if clear { // 避免暂时性的内存泄露。
		temp := s[len(result):]
		for i := range temp {
			temp[i] = t0 // t0是类型T的零值
		}
	}
	return result
}

func main() {
}
