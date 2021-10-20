package main

import "fmt"

func init() {
	// 结果为complex128类型的1.0+0.0i。虚部被舍入了。
	fmt.Println(complex128(1 + -1e-1000i))
	// 结果为float32类型的0.5。这里也舍入了。
	fmt.Println(float32(0.49999999))
	// 只要目标类型不是整数类型，舍入都是允许的。
	fmt.Println(float32(17000000000000000))
	fmt.Println(float32(123))
	fmt.Println(uint(1.0))
	fmt.Println(int8(-123))
	fmt.Println(int16(6+0i))
	fmt.Println(complex128(789))

	fmt.Println(string(65))          // "A"
	fmt.Println(string('A'))         // "A"
	fmt.Println(string('\u68ee'))    // "森"
	fmt.Println(string(-1))          // "\uFFFD"
	fmt.Println(string(0xFFFD))      // "\uFFFD"
	fmt.Println(string(0x2FFFFFFFF)) // "\uFFFD"
}

func init() {
	//fmt.Println(int(1.23))    // 1.23不能被表示为int类型值。
	//fmt.Println(uint8(-1))     // -1不能被表示为uint8类型值。
	//fmt.Println(float64(1+2i)) // 1+2i不能被表示为float64类型值。

	// -1e+1000不能被表示为float64类型值。不允许溢出。
	//fmt.Println(float64(-1e1000))
	// 0x10000000000000000做为int值将溢出。
	//fmt.Println(int(0x10000000000000000))

	// 字面量65.0的默认类型是float64（不是一个整数类型）。
	//fmt.Println(string(65.0))
	// 66+0i的默认类型是complex128（不是一个整数类型）。
	//fmt.Println(string(66+0i))

}

// Declare two individual constants. Yes,
// non-ASCII letters can be used in identifiers.
const π = 3.1416
const Pi = π // <=> const Pi = 3.1416

// Declare multiple constants in a group.
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "radian"
)

func init() {
	// Declare multiple constants in one line.
	const TwoPi, HalfPi, Unit2 = π * 2, π * 0.5, "degree"
}

func init() {
	const X float32 = 3.14

	const (
		A, B int64   = -3, 5
		Y    float32 = 2.718
	)
}

func init() {
	const X = float32(3.14)

	const (
		A, B = int64(-3), int64(5)
		Y    = float32(2.718)
	)
}

func init() {
	// error: 256 overflows uint8
	const a uint8 = 256
	// error: 256 overflows uint8
	const b = uint8(255) + uint8(1)
	// error: 128 overflows int8
	const c = int8(-128) / int8(-1)
	// error: -1 overflows uint
	const MaxUint_a = uint(^0)
	// error: -1 overflows uint
	const MaxUint_b uint = ^0
}

func init() {
	const MaxUint uint = (1 << 64) - 1
}

func init() {
	const MaxUint = ^uint(0)
}

func init() {
	const MaxInt = int(^uint(0) >> 1)
}

func init() {
	const NativeWordBits = 32 << (^uint(0) >> 63) // 64 or 32
	const Is64bitOS = ^uint(0) >> 63 != 0
	const Is32bitOS = ^uint(0) >> 32 == 0
}

func init() {
	const (
		X float32 = 3.14
		Y           // here must be one identifier
		Z           // here must be one identifier

		A, B = "Go", "language"
		C, _
		// In the above line, the blank identifier
		// is required to be present.
	)
}

func init() {
	const (
		X float32 = 3.14
		Y float32 = 3.14
		Z float32 = 3.14

		A, B = "Go", "language"
		C, _ = "Go", "language"
	)
}

func main() {}