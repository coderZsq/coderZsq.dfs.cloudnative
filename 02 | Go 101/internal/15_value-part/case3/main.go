package main

type _string struct {
	elements *byte // 引用着底层的byte元素
	len      int   // 字符串的长度
}
