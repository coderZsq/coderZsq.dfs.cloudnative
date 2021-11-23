package main

import "unsafe"

type _interface struct {
	dynamicType  *_type         // 引用着接口值的动态类型
	dynamicValue unsafe.Pointer // 引用着接口值的动态值
}

