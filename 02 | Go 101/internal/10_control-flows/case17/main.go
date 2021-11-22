package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(3); n {
	case 0: fmt.Println("n == 0")
	case 1: fmt.Println("n == 1")
	default: fmt.Println("n == 2")
	}

	switch n := rand.Intn(3); n {
	default: fmt.Println("n == 2")
	case 0: fmt.Println("n == 0")
	case 1: fmt.Println("n == 1")
	}

	switch n := rand.Intn(3); n {
	case 0: fmt.Println("n == 0")
	default: fmt.Println("n == 2")
	case 1: fmt.Println("n == 1")
	}
}