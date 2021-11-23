package main

type Book struct {
	name string
	// 更多其它字段……
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "我是一个自定义整数类型"
}

func main() {

}
