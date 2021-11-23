package main

import "fmt"

func main() {
	var pa *[5]int // == nil
	fmt.Println(len(pa), cap(pa)) // 5 5
}
