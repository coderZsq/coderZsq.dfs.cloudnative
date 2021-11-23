package main

import "unsafe"

type _slice struct {
	elements unsafe.Pointer // 引用着底层存储在间接部分上的元素
	len      int            // 长度
	cap      int            // 容量
}

func main() {

}
