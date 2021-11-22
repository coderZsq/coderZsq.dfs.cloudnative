package main

import "fmt"

func main() {
	switch { // <=> switch true {
	case true: fmt.Println("hello")
	default: fmt.Println("bye")
	}
}
