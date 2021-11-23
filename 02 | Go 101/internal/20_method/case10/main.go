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

func main() {
	var books = Books{{123}, {456}}
	books.Modify()
	fmt.Println(books) // [{500} {456}]
}
