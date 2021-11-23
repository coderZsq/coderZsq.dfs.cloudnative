package main

type LangCategory struct {
	dynamic bool
	strong  bool
}

var _ = map[LangCategory]map[string]int{
	{true, true}: {
		"Python": 1991,
		"Erlang": 1986,
	},
	{true, false}: {
		"JavaScript": 1995,
	},
	{false, true}: {
		"Go":   2009,
		"Rust": 2010,
	},
	{false, false}: {
		"C": 1972,
	},
}

func main() {

}