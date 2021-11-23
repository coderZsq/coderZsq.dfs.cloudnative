package main

import "fmt"

func main() {
	var s = append([]string(nil), "array", "slice")
	fmt.Println(s)      // [array slice]
	fmt.Println(cap(s)) // 2
	s = append(s, "map")
	fmt.Println(s)      // [array slice map]
	fmt.Println(cap(s)) // 4
	s = append(s, "channel")
	fmt.Println(s)      // [array slice map channel]
	fmt.Println(cap(s)) // 4
}
