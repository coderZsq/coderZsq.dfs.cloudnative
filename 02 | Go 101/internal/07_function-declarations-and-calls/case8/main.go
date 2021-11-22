package main

// c是一个类型不确定复数常量。
const c = complex(1.6, 3.3)

// 函数调用real(c)和imag(c)的结果都是类型
// 不确定浮点数值。在下面这句赋值中，它们都
// 被推断为float32类型的值。
var a, b float32 = real(c), imag(c)

// 变量d的类型被推断为内置类型complex64。
// 函数调用real(d)和imag(d)的结果都是
// 类型为float32的类型确定值。
var d = complex(a, b)

// 变量e的类型被推断为内置类型complex128。
// 函数调用real(e)和imag(e)的结果都是
// 类型为float64的类型确定值。
var e = c

func main() {

}