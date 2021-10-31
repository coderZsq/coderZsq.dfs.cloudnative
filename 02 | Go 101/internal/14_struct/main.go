package main

import "fmt"

func init() {
	var _ struct {
		title  string
		author string
		pages  int
	}

	var _ struct {
		title, author string
		pages         int
	}

	var _ struct {
		Title  string `json:"title" myfmt:"s1"`
		Author string `json:"author,omitempty" myfmt:"s2"`
		Pages  int    `json:"pages,omitempty" myfmt:"n1"`
		X, Y   bool   `myfmt:"b1"`
	}
}

type Book struct {
	title, author string
	pages         int
}

func init() {
	book := Book{"Go语言101", "老貘", 256}
	fmt.Println(book) // {Go语言101 老貘 256}

	// 使用带字段名的组合字面量来表示结构体值。
	book = Book{author: "老貘", pages: 256, title: "Go语言101"}
	// title和author字段的值都为空字符串""，pages字段的值为0。
	book = Book{}
	// title字段空字符串""，pages字段为0。
	book = Book{author: "老貘"}

	// 使用选择器来访问和修改字段值。
	var book2 Book // <=> book2 := Book{}
	book2.author = "Tapir"
	book2.pages = 300
	fmt.Println(book2.pages) // 300
}

func init() {
	var _ = Book{
		author: "老貘",
		pages:  256,
		title:  "Go语言101", // 这里行尾的逗号不可省略
	}

	// 下行}前的逗号可以省略。
	var _ = Book{author: "老貘", pages: 256, title: "Go语言101"}
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

func init() {
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

func init() {
	type Book struct {
		Pages int
	}
	// Book{100}是不可寻址的，但是它可以被取地址。
	p := &Book{100} // <=> tmp := Book{100}; p := &tmp
	p.Pages = 200
}

func init() {
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

type S0 struct {
	y int "foo"
	x bool
}

type S1 = struct { // S1是一个非定义类型
	x int "foo"
	y bool
}

type S2 = struct { // S2也是一个非定义类型
	x int "bar"
	y bool
}

type S3 S2 // S3是一个定义类型。
type S4 S3 // S4是一个定义类型。
// 如果不考虑字段标签，S3（S4）和S1的底层类型一样。
// 如果考虑字段标签，S3（S4）和S1的底层类型不一样。

var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}
func f2() {
	v1 = S1(v2); v2 = S2(v1)
	v1 = S1(v3); v3 = S3(v1)
	v1 = S1(v4); v4 = S4(v1)
	v2 = v3; v3 = v2 // 这两个转换可以是隐式的
	v2 = v4; v4 = v2 // 这两个转换也可以是隐式的
	v3 = S3(v4); v4 = S4(v3)
}

func init() {
	var aBook = struct {
		author struct { // 此字段的类型为一个匿名结构体类型
			firstName, lastName string
			gender              bool
		}
		title string
		pages int
	}{
		author: struct {
			firstName, lastName string
			gender              bool
		}{
			firstName: "Mark",
			lastName: "Twain",
		}, // 此组合字面量中的类型为一个匿名结构体类型
		title: "The Million Pound Note",
		pages: 96,
	}
	_ = aBook
}

func main() {

}
