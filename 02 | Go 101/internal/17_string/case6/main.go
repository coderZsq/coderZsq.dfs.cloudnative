package main

import "fmt"

func main() {
	var str = "world"
	// 这里，转换[]byte(str)将不需要一个深复制。
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	key := []byte{'k', 'e', 'y'}
	m := map[string]string{}
	// 这个string(key)转换仍然需要深复制。
	m[string(key)] = "value"
	// 这里的转换string(key)将不需要一个深复制。
	// 即使key是一个包级变量，此优化仍然有效。
	fmt.Println(m[string(key)]) // value
}
