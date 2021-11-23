package main

import "fmt"

func main() {
	type Book struct {
		Pages int
	}
	var book = Book{} // 变量值book是可寻址的
	p := &book.Pages
	*p = 123
	fmt.Println(book) // {123}

	// 下面这两行编译不通过，因为Book{}是不可寻址的，
	// 继而Book{}.Pages也是不可寻址的。
	/*
		Book{}.Pages = 123
		p = &Book{}.Pages // <=> p = &(Book{}.Pages)
	*/
}