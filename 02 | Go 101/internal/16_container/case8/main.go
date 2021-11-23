package main

type language struct {
	name string
	year int
}

var _ = [...]language{
	language{"C", 1972},
	language{"Python", 1991},
	language{"Go", 2009},
}

func main() {

}
