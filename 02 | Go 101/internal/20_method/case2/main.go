package main

import "fmt"

type Book struct {
	pages int
}

type Books []Book

func (books Books) Modify() {
	// 对属主参数的间接部分的修改将反映到方法之外。
	books[0].pages = 500
	// 对属主参数的直接部分的修改不会反映到方法之外。
	books = append(books, Book{789})
}

func init() {
	var books = Books{{123}, {456}}
	books.Modify()
	fmt.Println(books) // [{500} {456}]
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) Pages2() int {
	return (*b).Pages()
}

func init() {
	var b = Book{pages: 123}
	var p = &b
	var f1 = b.Pages
	var f2 = p.Pages
	var g1 = p.Pages2
	var g2 = b.Pages2
	b.pages = 789
	fmt.Println(f1()) // 123
	fmt.Println(f2()) // 123
	fmt.Println(g1()) // 789
	fmt.Println(g2()) // 789
}

func main() {

}
