package main

import "fmt"

func main() {
	s := "éक्षिaπ囧"
	for i, rn := range s {
		fmt.Printf("%2v: 0x%x %v \n", i, rn, string(rn))
	}
	fmt.Println(len(s))
}
