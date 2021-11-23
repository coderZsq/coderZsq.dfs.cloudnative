package main

type LangCategory struct {
	dynamic bool
	strong  bool
}

// 此映射值的类型的键值类型为一个结构体类型，
// 元素类型为另一个映射类型：map[string]int。
var _ = map[LangCategory]map[string]int{
	LangCategory{true, true}: map[string]int{
		"Python": 1991,
		"Erlang": 1986,
	},
	LangCategory{true, false}: map[string]int{
		"JavaScript": 1995,
	},
	LangCategory{false, true}: map[string]int{
		"Go":   2009,
		"Rust": 2010,
	},
	LangCategory{false, false}: map[string]int{
		"C": 1972,
	},
}

func main() {

}