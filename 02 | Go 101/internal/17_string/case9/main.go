package main

import "fmt"

func main() {
	s := "éक्षिaπ囧"
	for i := 0; i < len(s); i++ {
		fmt.Printf("第%v个字节为0x%x\n", i, s[i])
	}
}
