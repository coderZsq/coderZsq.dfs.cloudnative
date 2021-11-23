package main

import "fmt"

func main() {
	var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	for i, n := range p { // panic
		fmt.Println(i, n)
	}
}