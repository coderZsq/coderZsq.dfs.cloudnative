package main

func main() {
	// 结果为complex128类型的1.0+0.0i。虚部被舍入了。
	_ = complex128(1 + -1e-1000i)
	// 结果为float32类型的0.5。这里也舍入了。
	_ = float32(0.49999999)
	// 只要目标类型不是整数类型，舍入都是允许的。
	_ = float32(17000000000000000)
	_ = float32(123)
	_ = uint(1.0)
	_ = int8(-123)
	_ = int16(6+0i)
	_ = complex128(789)

	_ = string(65)          // "A"
	_ = string('A')         // "A"
	_ = string('\u68ee')    // "森"
	_ = string(-1)          // "\uFFFD"
	_ = string(0xFFFD)      // "\uFFFD"
	_ = string(0x2FFFFFFFF) // "\uFFFD"
}