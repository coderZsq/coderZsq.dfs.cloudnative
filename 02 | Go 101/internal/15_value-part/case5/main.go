package main

import "unsafe"

type _interface struct {
	dynamicTypeInfo *struct {
		dynamicType *_type       // 引用着接口值的动态类型
		methods     []*_function // 引用着动态类型的对应方法列表
	}
	dynamicValue unsafe.Pointer // 引用着动态值
}
