package main

func main() {
	_ = 1 + 2i       // == 1.0 + 2.0i
	_ = 1. - .1i     // == 1.0 + -0.1i
	_ = 1.23i - 7.89 // == -7.89 + 1.23i
	_ = 1.23i        // == 0.0 + 1.23i
}