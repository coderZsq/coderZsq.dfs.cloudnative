package main

func f() []int {
	s := make([]int, 10, 100)
	return s[50:60]
}