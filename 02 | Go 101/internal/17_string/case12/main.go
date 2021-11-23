package main

import (
	"fmt"
	"time"
)

func main() {
	bs := make([]byte, 1<<26)
	s0 := string(bs)
	s1 := string(bs)
	s2 := s1

	// s0、s1和s2是三个相等的字符串。
	// s0的底层字节序列是bs的一个深复制。
	// s1的底层字节序列也是bs的一个深复制。
	// s0和s1底层字节序列为两个不同的字节序列。
	// s2和s1共享同一个底层字节序列。

	startTime := time.Now()
	_ = s0 == s1
	duration := time.Now().Sub(startTime)
	fmt.Println("duration for (s0 == s1):", duration)

	startTime = time.Now()
	_ = s1 == s2
	duration = time.Now().Sub(startTime)
	fmt.Println("duration for (s1 == s2):", duration)
}
