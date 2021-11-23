package main

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
