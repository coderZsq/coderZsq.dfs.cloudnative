package main

import (
	"fmt"
	"time"
)

func main() {
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("现在为上午。")
	} else if h > 19 {
		fmt.Println("现在为晚上。")
	} else {
		fmt.Println("现在为下午。")
		// 左h是一个新声明的变量，右h已经在上面声明了。
		h := h
		// 刚声明的h遮掩了上面声明的h。
		_ = h
	}

	// 上面声明的两个h在此处都不可见。
}