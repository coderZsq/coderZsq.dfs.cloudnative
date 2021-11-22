package main

func main() {
	// 变量lang和year都为新声明的变量。
	lang, year := "Go language", 2007

	// 这里，只有变量createdBy是新声明的变量。
	// 变量year已经在上面声明过了，所以这里仅仅
	// 改变了它的值，或者说它被重新声明了。
	year, createdBy := 2009, "Google Research"

	// 这是一个纯赋值语句。
	lang, year = "Go", 2012

	print(lang, "由", createdBy, "发明")
	print("并发布于", year, "年。")
	println()
}
