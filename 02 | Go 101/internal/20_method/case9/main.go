package main

import "fmt"

type Book struct {
	pages int
}

func (b Book) SetPages(pages int) {
	b.pages = pages
}

func main() {
	var b Book
	b.SetPages(123)
	fmt.Println(b.pages) // 0
}
