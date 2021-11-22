package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if n := rand.Int(); n%2 == 0 {
		fmt.Println(n, "是一个偶数。")
	} else {
		fmt.Println(n, "是一个奇数。")
	}

	n := rand.Int() % 2 // 此n不是上面声明的n
	if n % 2 == 0 {
		fmt.Println("一个偶数。")
	}

	if ; n % 2 != 0 {
		fmt.Println("一个奇数。")
	}
}

