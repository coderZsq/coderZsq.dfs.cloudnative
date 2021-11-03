package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

type _string struct {
	elements *byte // 引用着底层的字节
	len      int   // 字符串中的字节数
}

func init() {
	const World = "world"
	var hello = "hello"

	// 衔接字符串。
	var helloWorld = hello + " " + World
	helloWorld += "!"
	fmt.Println(helloWorld) // hello world!

	// 比较字符串。
	fmt.Println(hello == "hello")   // true
	fmt.Println(hello > helloWorld) // false
}

func init() {
	var helloWorld = "hello world!"

	var hello = helloWorld[:5] // 取子字符串
	// 104是英文字符h的ASCII（和Unicode）码。
	fmt.Println(hello[0])         // 104
	fmt.Printf("%T \n", hello[0]) // uint8

	// hello[0]是不可寻址和不可修改的，所以下面
	// 两行编译不通过。
	/*
		hello[0] = 'H'         // error
		fmt.Println(&hello[0]) // error
	*/

	// 下一条语句将打印出：5 12 true
	fmt.Println(len(hello), len(helloWorld),
		strings.HasPrefix(helloWorld, hello))
}

func init() {
	const s = "Go101.org" // len(s) == 9

	// len(s)是一个常量表达式，但len(s[:])却不是。
	var a byte = 1 << len(s) / 128
	var b byte = 1 << len(s[:]) / 128
	fmt.Println(a, b) // 4 0
}

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}


func init() {
	s := "颜色感染是一个有趣的游戏。"
	bs := []byte(s) // string -> []byte
	s = string(bs)  // []byte -> string
	rs := []rune(s) // string -> []rune
	s = string(rs)  // []rune -> string
	rs = bytes.Runes(bs) // []byte -> []rune
	bs = Runes2Bytes(rs) // []rune -> []byte
}

func init() {
	var str = "world"
	// 这里，转换[]byte(str)将不需要一个深复制。
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	key := []byte{'k', 'e', 'y'}
	m := map[string]string{}
	// 这个string(key)转换仍然需要深复制。
	m[string(key)] = "value"
	// 这里的转换string(key)将不需要一个深复制。
	// 即使key是一个包级变量，此优化仍然有效。
	fmt.Println(m[string(key)]) // value
}


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

func init() {
	fmt.Println(testing.AllocsPerRun(1, fc)) // 1
	fmt.Println(testing.AllocsPerRun(1, fd)) // 3
}

func init() {
	s := "éक्षिaπ囧"
	for i, rn := range s {
		fmt.Printf("%2v: 0x%x %v \n", i, rn, string(rn))
	}
	fmt.Println(len(s))
}

func init() {
	s := "éक्षिaπ囧"
	for i := 0; i < len(s); i++ {
		fmt.Printf("第%v个字节为0x%x\n", i, s[i])
	}
}

func init() {
	s := "éक्षिaπ囧"
	// 这里，[]byte(s)不需要深复制底层字节。
	for i, b := range []byte(s) {
		fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	}
}

func init() {
	hello := []byte("Hello ")
	world := "world!"

	// helloWorld := append(hello, []byte(world)...) // 正常的语法
	helloWorld := append(hello, world...)            // 语法糖
	fmt.Println(string(helloWorld))

	helloWorld2 := make([]byte, len(hello) + len(world))
	copy(helloWorld2, hello)
	// copy(helloWorld2[len(hello):], []byte(world)) // 正常的语法
	copy(helloWorld2[len(hello):], world)            // 语法糖
	fmt.Println(string(helloWorld2))
}

func init() {
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

func main() {

}