package main

import (
	format "fmt"
	random "math/rand"
	"time"
)

func main() {
	random.Seed(time.Now().UnixNano())
	format.Print("一个随机数:", random.Uint32(), "\n")

	// 下面这两行编译不通过，因为rand不可识别。
	/*
		rand.Seed(time.Now().UnixNano())
		fmt.Print("一个随机数:", rand.Uint32(), "\n")
	*/
}