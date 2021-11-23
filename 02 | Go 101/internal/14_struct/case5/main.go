package main

type Book struct {
	title, author string
	pages         int
}

func main() {
	var _ = Book {
		author: "老貘",
		pages: 256,
		title: "Go语言101", // 这里行尾的逗号不可省略
	}

	// 下行}前的逗号可以省略。
	var _ = Book{author: "老貘", pages: 256, title: "Go语言101",}
}
