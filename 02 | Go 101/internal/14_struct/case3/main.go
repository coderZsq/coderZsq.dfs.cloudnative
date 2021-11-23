package main

func main() {
	var _ struct {
		Title  string `json:"title" myfmt:"s1"`
		Author string `json:"author,omitempty" myfmt:"s2"`
		Pages  int    `json:"pages,omitempty" myfmt:"n1"`
		X, Y   bool   `myfmt:"b1"`
	}
}
