package main

import "fmt"

/* Some type definition declarations */

// status and bool are two different types.
type status bool
// MyString and string are two different types.
type MyString string
// Id and uint64 are two different types.
type Id uint64
// real and float32 are two different types.
type real float32

/* Some type alias declarations */

// boolean and bool denote the same type.
type boolean = bool
// Text and string denote the same type.
type Text = string
// U8, uint8 and byte denote the same type.
type U8 = uint8
// char, rune and int32 denote the same type.
type char = rune

// Integer value literals
func init() {
	fmt.Println(0xF) // the hex form (starts with a "0x" or "0X")
	fmt.Println(0XF)

	fmt.Println(017) // the octal form (starts with a "0", "0o" or "0O")
	fmt.Println(0o17)
	fmt.Println(0O17)

	fmt.Println(0b1111) // the binary form (starts with a "0b" or "0B")
	fmt.Println(0B1111)

	fmt.Println(15) // the decimal form (starts without a "0")

	println(15 == 017) // true
	println(15 == 0xF) // true
}

// Floating-point value literals
func init() {
	fmt.Println(1.23)
	fmt.Println(01.23) // == 1.23
	fmt.Println(.23)
	fmt.Println(1.)
	// An "e" or "E" starts the exponent part (10-based).
	fmt.Println(1.23e2)  // == 123.0
	fmt.Println(123E2)   // == 12300.0
	fmt.Println(123.E+2) // == 12300.0
	fmt.Println(1e-1)    // == 0.1
	fmt.Println(.1e0)    // == 0.1
	fmt.Println(0010e-2) // == 0.1
	fmt.Println(0e+5)    // == 0.0

	fmt.Println(0x1p-2)     // == 1.0/4 = 0.25
	fmt.Println(0x2.p10)    // == 2.0 * 1024 == 2048.0
	fmt.Println(0x1.Fp+0)   // == 1+15.0/16 == 1.9375
	fmt.Println(0X.8p1)     // == 8.0/16 * 2 == 1.0
	fmt.Println(0X1FFFP-16) // == 0.1249847412109375

	//fmt.Println(0x.p1)    // mantissa has no digits
	//fmt.Println(1p-2)     // p exponent requires hexadecimal mantissa
	//fmt.Println(0x1.5e-2) // hexadecimal mantissa requires p exponent

	fmt.Println(0x15e-2) // == 0x15e - 2 // a subtraction expression
}

// Imaginary value literals
func init() {
	fmt.Println(1.23i)
	fmt.Println(1.i)
	fmt.Println(.23i)
	fmt.Println(123i)
	fmt.Println(0123i)   // == 123i (for backward-compatibility. See below.)
	fmt.Println(1.23E2i) // == 123i
	fmt.Println(1e-1i)
	fmt.Println(011i)   // == 11i (for backward-compatibility. See below.)
	fmt.Println(00011i) // == 11i (for backward-compatibility. See below.)
	// The following lines only compile okay since Go 1.13.
	fmt.Println(0o11i)    // == 9i
	fmt.Println(0x11i)    // == 17i
	fmt.Println(0b11i)    // == 3i
	fmt.Println(0X.8p-0i) // == 0.5i

	fmt.Println(1 + 2i)       // == 1.0 + 2.0i
	fmt.Println(1. - .1i)     // == 1.0 + -0.1i
	fmt.Println(1.23i - 7.89) // == -7.89 + 1.23i
	fmt.Println(1.23i)        // == 0.0 + 1.23i

	// Legal ones:
	fmt.Println(6_9)          // == 69
	fmt.Println(0_33_77_22)   // == 0337722
	fmt.Println(0x_Bad_Face)  // == 0xBadFace
	fmt.Println(0X_1F_FFP-16) // == 0X1FFFP-16
	fmt.Println(0b1011_0111 + 0xA_B.Fp2i)
	
	// Illegal ones:
	//fmt.Println(_69)        // _ can't appear as the first character
	//fmt.Println(69_)        // _ can't appear as the last character
	//fmt.Println(6__9)       // one side of _ is a illegal character
	//fmt.Println(0_xBadFace) // "x" is not a legal octal digit
	//fmt.Println(1_.5)       // "." is not a legal octal digit
	//fmt.Println(1._5)       // "." is not a legal octal digit
}

// Rune value literals
func init() {
	fmt.Println('a') // an English character
	fmt.Println('π')
	fmt.Println('众') // a Chinese character

	// 141 is the octal representation of decimal number 97.
	fmt.Println('\141')
	// 61 is the hex representation of decimal number 97.
	fmt.Println('\x61')
	fmt.Println('\u0061')
	fmt.Println('\U00000061')

	println('a' == 97)
	println('a' == '\141')
	println('a' == '\x61')
	println('a' == '\u0061')
	println('a' == '\U00000061')
	println(0x61 == '\x61')
	println('\u4f17' == '众')

	//\a   (Unicode value 0x07) alert or bell
	//\b   (Unicode value 0x08) backspace
	//\f   (Unicode value 0x0C) form feed
	//\n   (Unicode value 0x0A) line feed or newline
	//\r   (Unicode value 0x0D) carriage return
	//\t   (Unicode value 0x09) horizontal tab
	//\v   (Unicode value 0x0b) vertical tab
	//\\   (Unicode value 0x5c) backslash
	//\'   (Unicode value 0x27) single quote

	println('\n') // 10
	println('\r') // 13
	println('\'') // 39

	println('\n' == 10)     // true
	println('\n' == '\x0A') // true
}

// String value literals
func init() {
	// The interpreted form.
	fmt.Println("Hello\nworld!\n\"你好世界\"")
	
	// The raw form.
	fmt.Println(`Hello
	world!
	"你好世界"`)

	// The following interpreted string literals are equivalent.
	fmt.Println("\141\142\143")
	fmt.Println("\x61\x62\x63")
	fmt.Println("\x61b\x63")
	fmt.Println("abc")

	// The following interpreted string literals are equivalent.
	fmt.Println("\u4f17\xe4\xba\xba")
	// The Unicode of 众 is 4f17, which is
	// UTF-8 encoded as three bytes: e4 bc 97.
	fmt.Println("\xe4\xbc\x97\u4eba")
	// The Unicode of 人 is 4eba, which is
	// UTF-8 encoded as three bytes: e4 ba ba.
	fmt.Println("\xe4\xbc\x97\xe4\xba\xba")
	fmt.Println("众人")
}

func main() {}