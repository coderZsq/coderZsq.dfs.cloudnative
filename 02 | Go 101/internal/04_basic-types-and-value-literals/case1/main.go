package main

// 一些类型定义声明
type status bool     // status和bool是两个不同的类型
type MyString string // MyString和string是两个不同的类型
type Id uint64       // Id和uint64是两个不同的类型
type real float32    // real和float32是两个不同的类型

// 一些类型别名声明
type boolean = bool // boolean和bool表示同一个类型
type Text = string  // Text和string表示同一个类型
type U8 = uint8     // U8、uint8和 byte表示同一个类型
type char = rune    // char、rune和int32表示同一个类型