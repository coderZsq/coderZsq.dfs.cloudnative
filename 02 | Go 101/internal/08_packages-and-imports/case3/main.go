package main

// 一条包引入语句引入了三个代码包。
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	fmt.Printf("下一个伪随机数是%v。\n", rand.Uint32())
}