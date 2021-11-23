package main

import "unsafe"

type _slice struct {
	elements unsafe.Pointer // 引用着底层的元素
	len      int            // 当前的元素个数
	cap      int            // 切片的容量
}
