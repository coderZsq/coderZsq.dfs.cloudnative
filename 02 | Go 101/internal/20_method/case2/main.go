package main

func Book.Pages(b Book) int {
	return b.pages // 此函数体和Book类型的Pages方法体一样
}

func (*Book).SetPages(b *Book, pages int) {
	b.pages = pages // 此函数体和*Book类型的SetPages方法体一样
}