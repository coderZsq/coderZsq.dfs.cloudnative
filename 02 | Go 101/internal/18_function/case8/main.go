package main

import "fmt"

func Concat(sep string, tokens ...string) (r string) {
	for i, t := range tokens {
		if i != 0 {
			r += sep
		}
		r += t
	}
	return
}

func main() {
	tokens := []string{"Go", "C", "Rust"}
	langsA := Concat(",", tokens...)        // 风格1
	langsB := Concat(",", "Go", "C","Rust") // 风格2
	fmt.Println(langsA == langsB)           // true
}

