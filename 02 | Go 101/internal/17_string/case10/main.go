package main

import "fmt"

func main() {
	s := "éक्षिaπ囧"
	// 这里，[]byte(s)不需要深复制底层字节。
	for i, b := range []byte(s) {
		fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	}
}
