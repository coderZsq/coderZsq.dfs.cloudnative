package main

func main() {
	println('a' == 97)
	println('a' == '\141')
	println('a' == '\x61')
	println('a' == '\u0061')
	println('a' == '\U00000061')
	println(0x61 == '\x61')
	println('\u4f17' == '众')
}