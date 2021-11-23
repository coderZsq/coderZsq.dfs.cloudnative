package main

import "fmt"

func main() {
	langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScript": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}
	// 此映射的键值和元素类型均为指针类型。
	// 这有些不寻常，只是为了讲解目的。
	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range langs {
		m0[&category] = &langInfo
		// 下面这行修改对映射langs没有任何影响。
		category.dynamic, category.strong = true, true
	}
	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	m1 := map[struct{ dynamic, strong bool }]map[string]int{}
	for category, langInfo := range m0 {
		m1[*category] = *langInfo
	}
	// 映射m0和m1中均只有一个条目。
	fmt.Println(len(m0), len(m1)) // 1 1
	fmt.Println(m1) // map[{true true}:map[C:1972]]
}
