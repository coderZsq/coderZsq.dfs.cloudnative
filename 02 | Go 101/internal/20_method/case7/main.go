package main

import "fmt"

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func main() {
	var book Book

	fmt.Printf("%T \n", book.Pages)       // func() int
	fmt.Printf("%T \n", (&book).SetPages) // func(int)
	// &book值有一个隐式方法Pages。
	fmt.Printf("%T \n", (&book).Pages)    // func() int

	// 调用这三个方法。
	(&book).SetPages(123)
	book.SetPages(123)           // 等价于上一行
	fmt.Println(book.Pages())    // 123
	fmt.Println((&book).Pages()) // 123
}
