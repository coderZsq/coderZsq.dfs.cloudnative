package main

func main() {
	type Book struct {
		Pages int
	}
	// Book{100}是不可寻址的，但是它可以被取地址。
	p := &Book{100} // <=> tmp := Book{100}; p := &tmp
	p.Pages = 200
}

