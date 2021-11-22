package main

type MyInt int
func (mi MyInt) IsOdd() bool {
	return mi%2 == 1
}

type Age MyInt

type X struct {
	MyInt
}
func (x X) Double() MyInt {
	return x.MyInt + x.MyInt
}

type Y struct {
	Age
}

type Z X