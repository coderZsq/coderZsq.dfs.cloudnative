package main

func main() {
	_ = '\141'   // 141是97的八进制表示
	_ = '\x61'   // 61是97的十六进制表示
	_ = '\u0061'
	_ = '\U00000061'
}