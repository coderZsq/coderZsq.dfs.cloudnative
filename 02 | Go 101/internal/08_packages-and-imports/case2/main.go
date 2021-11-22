package main

import "fmt"
import "math/rand"

func main() {
	fmt.Printf("下一个伪随机数总是%v。\n", rand.Uint32())
}
