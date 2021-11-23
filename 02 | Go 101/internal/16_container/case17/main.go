package main

import "fmt"

func main() {
	m := map[string]int{"Go": 2007}
	m["C"] = 1972     // 添加
	m["Java"] = 1995  // 添加
	fmt.Println(m)    // map[C:1972 Go:2007 Java:1995]
	m["Go"] = 2009    // 修改
	delete(m, "Java") // 删除
	fmt.Println(m)    // map[C:1972 Go:2009]
}
