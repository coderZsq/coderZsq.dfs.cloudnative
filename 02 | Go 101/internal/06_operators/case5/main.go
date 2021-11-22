package main

const n = uint(2)
var m = uint(2)

// 这两行编译没问题。
var _ float64 = 1 << n
var _ = float64(1 << n)

// 这两行编译失败。
var _ float64 = 1 << m  // error
var _ = float64(1 << m) // error

//var _ = float64(1) << m
var _ = 1.0 << m // error: shift of type float64

func main() {

}