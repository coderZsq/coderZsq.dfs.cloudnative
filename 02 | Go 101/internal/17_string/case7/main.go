package main

import (
	"fmt"
	"testing"
)

var s string
var x = []byte{1023: 'x'}
var y = []byte{1023: 'y'}

func fc() {
	// 下面的四个转换都不需要深复制。
	if string(x) != string(y) {
		s = (" " + string(x) + string(y))[1:]
	}
}

func fd() {
	// 两个在比较表达式中的转换不需要深复制，
	// 但两个字符串衔接中的转换仍需要深复制。
	// 请注意此字符串衔接和fc中的衔接的差别。
	if string(x) != string(y) {
		s = string(x) + string(y)
	}
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, fc)) // 1
	fmt.Println(testing.AllocsPerRun(1, fd)) // 3
}
