package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := make([]int, 2, 6)
	fmt.Println(len(s), cap(s)) // 2 6

	reflect.ValueOf(&s).Elem().SetLen(3)
	fmt.Println(len(s), cap(s)) // 3 6

	reflect.ValueOf(&s).Elem().SetCap(5)
	fmt.Println(len(s), cap(s)) // 3 5
}
