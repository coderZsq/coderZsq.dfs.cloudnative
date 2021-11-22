package main

func main() {
	println('\n') // 10
	println('\r') // 13
	println('\'') // 39

	println('\n' == 10)     // true
	println('\n' == '\x0A') // true
}