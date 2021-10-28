package main

//import importname "path/to/package"

import fmt "fmt"        // <=> import "fmt"
import rand "math/rand" // <=> import "math/rand"
import time "time"      // <=> import "time"

//import (
//	"fmt"
//	"math/rand"
//	"time"
//)

func init() {
	fmt.Println("Go has", 25, "keywords.")
}

func init() {
	fmt.Printf("下一个伪随机数总是%v。\n", rand.Uint32())
}

func init() {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	fmt.Printf("下一个伪随机数是%v。\n", rand.Uint32())
}

func init() {
	a, b := 123, "Go"
	fmt.Printf("a == %v == 0x%x, b == %s\n", a, a, b)
	fmt.Printf("type of a: %T, type of b: %T\n", a, b)
}

func init() {
	fmt.Println("hi,", bob)
}

func main() {
	fmt.Println("bye")
}

func init() {
	fmt.Println("hello,", smith)
}

func titledName(who string) string {
	return "Mr. " + who
}

var bob, smith = titledName("Bob"), titledName("Smith")

func f() int {
	return z + y
}

func g() int {
	return y / 2
}

var (
	w       = x
	x, y, z = f(), 123, g()
)
