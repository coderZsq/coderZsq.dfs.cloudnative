package main

func main() {
	// 这几个字符串字面量是等价的。
	_ = "\141\142\143"
	_ = "\x61\x62\x63"
	_ = "\x61b\x63"
	_ = "abc"

	// 这几个字符串字面量是等价的。
	_ = "\u4f17\xe4\xba\xba"
	// “众”的Unicode值为4f17，它的UTF-8
	// 编码为三个字节：0xe4 0xbc 0x97。
	_ = "\xe4\xbc\x97\u4eba"
	// “人”的Unicode值为4eba，它的UTF-8
	// 编码为三个字节：0xe4 0xba 0xba。
	_ = "\xe4\xbc\x97\xe4\xba\xba"
	_ = "众人"
}