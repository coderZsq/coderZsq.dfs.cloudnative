package main

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return Book.Pages(b)
}

func (b *Book) SetPages(pages int) {
	(*Book).SetPages(b, pages)
}

func main() {

}