package main

import "unsafe"

func init() {
	type hashtableImpl struct {
	}

	// 映射类型
	type _map *hashtableImpl // 目前，官方标准编译器是使用
	// 哈希表来实现映射的。

	type channelImpl struct {
	}

	// 通道类型
	type _channel *channelImpl

	type functionImpl struct {
	}

	// 函数类型
	type _function *functionImpl
}

func init() {
	type _slice struct {
		elements unsafe.Pointer // 引用着底层的元素
		len      int            // 当前的元素个数
		cap      int            // 切片的容量
	}
}

func init() {
	type _string struct {
		elements *byte // 引用着底层的byte元素
		len      int   // 字符串的长度
	}

}

func init() {
	type _type struct {
	}

	type _interface struct {
		dynamicType  *_type         // 引用着接口值的动态类型
		dynamicValue unsafe.Pointer // 引用着接口值的动态值
	}
}

func init() {

	type _type struct {
	}

	type _function struct {
	}

	type _interface struct {
		dynamicTypeInfo *struct {
			dynamicType *_type       // 引用着接口值的动态类型
			methods     []*_function // 引用着动态类型的对应方法列表
		}
		dynamicValue unsafe.Pointer // 引用着动态值
	}
}
