package main

// 结果为非常量
var a, b uint8 = 255, 1
var c = a + b  // c==0。a+b是一个非常量表达式，
// 结果中溢出的高位比特将被截断舍弃。
var d = a << b // d == 254。同样，结果中溢出的
// 高位比特将被截断舍弃。

// 结果为类型不确定常量，允许溢出其默认类型。
const X = 0x1FFFFFFFF * 0x1FFFFFFFF // 没问题，尽管X溢出
const R = 'a' + 0x7FFFFFFF          // 没问题，尽管R溢出

// 运算结果或者转换结果为类型确定常量
var e = X                // error: X溢出int。
var h = R                // error: R溢出rune。
//const Y = 128 - int8(1)  // error: 128溢出int8。
const Z = uint8(255) + 1 // error: 256溢出uint8。

func main() {

}