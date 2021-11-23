package main

func main() {
	type Book struct {
		pages int
	}
	book1 := &Book{100} // book1是一个指针
	book2 := new(Book)  // book2是另外一个指针
	// 像使用结构值一样来使用结构体值的指针。
	book2.pages = book1.pages
	// 上一行等价于下一行。换句话说，上一行
	// 两个选择器中的指针属主将被自动解引用。
	(*book2).pages = (*book1).pages
}
