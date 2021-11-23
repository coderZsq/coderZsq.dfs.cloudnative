package main

// 注意：这不是合法的Go语法。这里这样表示只是
// 为了解释目的。它表明表达式(&aBook).Pages
// 将被估值为aBook.Pages（见随后几节）。
func (b *Book) Pages = (*b).Pages

func (*Book).Pages(b *Book) int {
	return Book.Pages(*b)
}
