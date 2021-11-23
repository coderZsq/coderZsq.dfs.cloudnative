package main

import "fmt"

func main() {
	s0 := []int{2, 3, 5}
	fmt.Println(s0, cap(s0)) // [2 3 5] 3
	s1 := append(s0, 7)      // 添加一个元素
	fmt.Println(s1, cap(s1)) // [2 3 5 7] 6
	s2 := append(s1, 11, 13) // 添加两个元素
	fmt.Println(s2, cap(s2)) // [2 3 5 7 11 13] 6
	s3 := append(s0)         // <=> s3 := s0
	fmt.Println(s3, cap(s3)) // [2 3 5] 3
	s4 := append(s0, s0...)  // 以s0为基础添加s0中所有的元素
	fmt.Println(s4, cap(s4)) // [2 3 5 2 3 5] 6

	s0[0], s1[0] = 99, 789
	fmt.Println(s2[0], s3[0], s4[0]) // 789 99 2
}
