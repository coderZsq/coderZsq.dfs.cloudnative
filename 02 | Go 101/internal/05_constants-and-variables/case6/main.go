package main

const a uint8 = 256             // error: 256溢出uint8
const b = uint8(255) + uint8(1) // error: 256溢出uint8
const c = int8(-128) / int8(-1) // error: 128溢出int8
const MaxUint_a = uint(^0)      // error: -1溢出uint
const MaxUint_b uint = ^0       // error: -1溢出uint

func main() {

}