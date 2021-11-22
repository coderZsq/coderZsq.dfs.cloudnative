package main

const N = 2
// A == 12，它是一个默认类型为int的类型不确定值。
const A = 3.0 << N
// B == 12，它是一个类型为int8的类型确定值。
const B = int8(3.0) << N

var m = uint(32)
// 下面的三行是相互等价的。
var x int64 = 1 << m  // 1的类型将被设想为int64，而非int
var y = int64(1 << m) // 同上
var z = int64(1) << m // 同上

// 下面这行编译不通过。
/*
var _ = 1.23 << m // error: 浮点数不能被移位
*/
