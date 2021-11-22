package foo

import "fmt"

type A struct {
	n int
}

func (a A) m() {
	fmt.Println("A", a.n)
}

type I interface {
	m()
}

func Bar(i I) {
	i.m()
}
