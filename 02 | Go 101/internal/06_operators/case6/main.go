package main

const n = uint(8)
var m = uint(8)

func main() {
	println(a, b) // 2 0
}

var a byte = 1 << n / 128
var b byte = 1 << m / 128

var _ = byte(int(1) << n / 128)
var _ = byte(1) << m / 128