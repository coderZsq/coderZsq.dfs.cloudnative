package main

// 映射类型
type _map *hashtableImpl // 目前，官方标准编译器是使用
// 哈希表来实现映射的。

// 通道类型
type _channel *channelImpl

// 函数类型
type _function *functionImpl
