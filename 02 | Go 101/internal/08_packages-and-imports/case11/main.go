package main

import (
	format "fmt"  // okay: 下面被使用了一次
	_ "math/rand" // okay: 匿名引入
)

func main() {
	format.Println() // 使用"fmt"包
}
