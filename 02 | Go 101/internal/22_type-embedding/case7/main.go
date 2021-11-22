package main

type I interface {
	m()
}

type T struct {
	I
}

func main() {
	var t T
	var i = &t
	t.I = i
	i.m() // 将调用t.m()，然后再次调用i.m()，......
}
