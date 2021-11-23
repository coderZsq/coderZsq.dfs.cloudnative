package main

import "fmt"

type Book struct {
	pages int
}

type Books []Book

func (books Books) Modify() {
	books = append(books, Book{789})
	books[0].pages = 500
}

func main() {
	var books = Books{{123}, {456}}
	books.Modify()
	fmt.Println(books) // [{123} {456}]
}
