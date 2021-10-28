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
	fmt.Println(int16(6 + 0i))
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
	//const a uint8 = 256
	// error: 256 overflows uint8
	//const b = uint8(255) + uint8(1)
	// error: 128 overflows int8
	//const c = int8(-128) / int8(-1)
	// error: -1 overflows uint
	//const MaxUint_a = uint(^0)
	// error: -1 overflows uint
	//const MaxUint_b uint = ^0
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
	const Is64bitOS = ^uint(0)>>63 != 0
	const Is32bitOS = ^uint(0)>>32 == 0
}

func init() {
	const (
		X float32 = 3.14
		Y         // here must be one identifier
		Z         // here must be one identifier

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

func init() {
	const iota = 0
}

func init() {
	const (
		k = 3 // now, iota == 0

		m float32 = iota + .5 // m float32 = 1 + .5
		n                     // n float32 = 2 + .5

		p    = 9          // now, iota == 3
		q    = iota * 2   // q = 4 * 2
		_                 // _ = 5 * 2
		r                 // r = 6 * 2
		s, t = iota, iota // s, t = 7, 7
		u, v              // u, v = 8, 8
		_, w              // _, w = 9, 9
	)

	const x = iota // x = 0
	const (
		y = iota // y = 0
		z        // z = 1
	)

	println(m)             // +1.500000e+000
	println(n)             // +2.500000e+000
	println(q, r)          // 8 12
	println(s, t, u, v, w) // 7 7 8 8 9
	println(x, y, z)       // 0 0 1
}

func init() {
	const (
		Failed    = iota - 1 // == -1
		Unknown              // == 0
		Succeeded            // == 1
	)

	const (
		Readable   = 1 << iota // == 1
		Writable               // == 2
		Executable             // == 4
	)
}

func init() {
	var lang, website string = "Go", "https://golang.org"
	var compiled, dynamic bool = true, false
	var announceYear int = 2009
	fmt.Println(lang, website)
	fmt.Println(compiled, dynamic)
	fmt.Println(announceYear)
}

func init() {
	// The types of the lang and dynamic variables
	// will be deduced as built-in types "string"
	// and "bool" by compilers, respectively.
	var lang, dynamic = "Go", false
	fmt.Println(lang, dynamic)

	// The types of the compiled and announceYear
	// variables will be deduced as built-in
	// types "bool" and "int", respectively.
	var compiled, announceYear = true, 2009
	fmt.Println(compiled, announceYear)

	// The types of the website variable will be
	// deduced as the built-in type "string".
	var website = "https://golang.org"
	fmt.Println(website)
}

func init() {
	// Both are initialized as blank strings.
	var lang, website string
	fmt.Println(lang, website)
	// Both are initialized as false.
	var interpreted, dynamic bool
	fmt.Println(interpreted, dynamic)
	// n is initialized as 0.
	var n int
	fmt.Println(n)
}

func init() {
	var (
		lang, bornYear, compiled     = "Go", 2007, true
		announceAt, releaseAt    int = 2009, 2012
		createdBy, website       string
	)
	fmt.Println(lang, bornYear, compiled)
	fmt.Println(announceAt, releaseAt)
	fmt.Println(createdBy, website)
}

func init() {
	const N = 123
	var x int
	var y, z float32

	//N = 9 // error: constant N is not modifiable
	y = N // ok: N is deduced as a float32 value
	//x = y // error: type mismatch
	x = N // ok: N is deduced as an int value
	//y = x // error: type mismatch
	z = y // ok
	_ = y // ok

	//x, y = y, x // error: type mismatch
	x, y = int(y), float32(x) // ok
	z, y = y, z               // ok
	_, y = y, z               // ok
	z, _ = y, z               // ok
	_, _ = y, z               // ok
	x, y = 69, 1.23           // ok
}

func init() {
	//var a, b int
	//a = b = 123 // syntax error
}

func init() {
	// Both lang and year are newly declared.
	lang, year := "Go language", 2007

	// Only createdBy is a new declared variable.
	// The year variable has already been
	// declared before, so here its value is just
	// modified, or we can say it is redeclared.
	year, createdBy := 2009, "Google Research"

	// This is a pure assignment.
	lang, year = "Go", 2012

	print(lang, " is created by ", createdBy)
	println(", and released at year", year)
}

// Some package-level variables.
var x, y, z = 123, true, "foo"

func init() {
	var q, r = 789, false
	r, s := true, "bar"
	r = y // r is unused.
	x = q // q is used.
	fmt.Println(r, s)
}

func init() {
	var q, r = 789, false
	r, s := true, "bar"
	r = y
	x = q

	_, _ = r, s // make r and s used.
}

var x1, y1 = a + 1, 5          // 8 5
var a, b, c = b + 1, c + 1, y1 // 7 6 5
func init() {
	fmt.Println(x1, y1, a, b, c)
}

//var x, y = y, x

func init() {
	const a = -1.23
	// The type of b is deduced as float64.
	var b = a
	// error: constant 1.23 truncated to integer.
	//var x = int32(a)
	// error: cannot assign float64 to int32.
	//var y int32 = b
	// okay: z == -1, and the type of z is int32.
	//       The fraction part of b is discarded.
	var z = int32(b)

	const k int16 = 255
	// The type of n is deduced as int16.
	var n = k
	// error: constant 256 overflows uint8.
	//var f = uint8(k + 1)
	// error: cannot assign int16 to uint8.
	//var g uint8 = n + 1
	// okay: h == 0, and the type of h is uint8.
	//       n+1 overflows uint8 and is truncated.
	var h = uint8(n + 1)

	_, _ = z, h
}

const y2 = 789

var x2 int = 123

func init() { // The x variable shadows the above declared
	// package-level variable x.
	var x2 = true

	// A nested code block.
	{
		// Here, the left x and y are both
		// new declared variable. The right
		// ones are declared in outer blocks.
		x2, y2 := x2, y2

		// In this code block, the just new
		// declared x and y shadow the outer
		// declared same-name identifiers.
		x2, z2 := !x2, y2/10 // only z is new declared
		y2 /= 100
		println(x2, y2, z2) // false 7 78
	}
	println(x) // true
	//println(z2) // error: z is undefined.
}

func init() {
	// 3 untyped named constants. Their bound
	// values all overflow their respective
	// default types. This is allowed.
	const n = 1 << 64          // overflows int
	const r = 'a' + 0x7FFFFFFF // overflows rune
	const x = 2e+308           // overflows float64

	_ = n >> 2
	_ = r - 0x7FFFFFFF
	_ = x / 2
}

func init() {
	// 3 typed named constants. Their bound
	// values are not allowed to overflow their
	// respective default types. The 3 lines
	// all fail to compile.
	//const n int = 1 << 64           // overflows int
	//const r rune = 'a' + 0x7FFFFFFF // overflows rune
	//const x float64 = 2e+308        // overflows float64
}

func init() {
	const X = 3
	const Y = X + X
	var a = X

	b := Y
	println(a, b, X, Y)
}

func init() {
	var a = 3

	b := 6
	println(a, b, 3, 6)
}

func main() {}
