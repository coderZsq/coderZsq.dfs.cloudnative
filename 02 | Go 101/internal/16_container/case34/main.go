package main

import "fmt"

func main() {
	m := map[string]int{"C": 1972, "C++": 1983, "Go": 2009}
	for lang, year := range m {
		fmt.Printf("%v: %v \n", lang, year)
	}

	a := [...]int{2, 3, 5, 7, 11}
	for i, prime := range a {
		fmt.Printf("%v: %v \n", i, prime)
	}

	s := []string{"go", "defer", "goto", "var"}
	for i, keyword := range s {
		fmt.Printf("%v: %v \n", i, keyword)
	}
}
