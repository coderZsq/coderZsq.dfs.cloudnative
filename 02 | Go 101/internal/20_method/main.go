package main

import "fmt"

// Age和int是两个不同的类型。我们不能为int和*int
// 类型声明方法，但是可以为Age和*Age类型声明方法。
type Age int
func (age Age) LargerThan(a Age) bool {
	return age > a
}
func (age *Age) Increase() {
	*age++
}

// 为自定义的函数类型FilterFunc声明方法。
type FilterFunc func(in int) bool
func (ff FilterFunc) Filter(in int) bool {
	return ff(in)
}

// 为自定义的映射类型StringSet声明方法。
type StringSet map[string]struct{}
func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// 为自定义的结构体类型Book和它的指针类型*Book声明方法。
type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

//func Book.Pages(b Book) int {
//	return b.pages // 此函数体和Book类型的Pages方法体一样
//}
//
//func (*Book).SetPages(b *Book, pages int) {
//	b.pages = pages // 此函数体和*Book类型的SetPages方法体一样
//}

func init() {
	var book Book
	// 调用这两个隐式声明的函数。
	(*Book).SetPages(&book, 123)
	fmt.Println(Book.Pages(book)) // 123
}

//func (b Book) Pages() int {
//	return Book.Pages(b)
//}
//
//func (b *Book) SetPages(pages int) {
//	(*Book).SetPages(b, pages)
//}

//// 注意：这不是合法的Go语法。这里这样表示只是
//// 为了解释目的。它表明表达式(&aBook).Pages
//// 将被估值为aBook.Pages（见随后几节）。
//func (b *Book) Pages = (*b).Pages

//func (*Book).Pages(b *Book) int {
//	return Book.Pages(*b)
//}

