package main

type Book struct {
	title, author string
	pages         int
}

func f() {
	book1 := Book{pages: 300}
	book2 := Book{"Go语言101", "老貘", 256}

	book2 = book1
	// 上面这行和下面这三行是等价的。
	book2.title = book1.title
	book2.author = book1.author
	book2.pages = book1.pages
}

func main() {

}