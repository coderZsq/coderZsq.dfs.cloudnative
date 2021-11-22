package main

import "fmt"

func main() {
	var i = 0
	for ; i < 10; {
		fmt.Println(i)
		i++
	}
	for i < 20 {
		fmt.Println(i)
		i++
	}
}