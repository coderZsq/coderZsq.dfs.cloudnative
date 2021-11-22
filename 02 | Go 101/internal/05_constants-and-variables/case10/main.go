package main

const NativeWordBits = 32 << (^uint(0) >> 63) // 64 or 32
const Is64bitOS = ^uint(0) >> 63 != 0
const Is32bitOS = ^uint(0) >> 32 == 0

func main() {

}
