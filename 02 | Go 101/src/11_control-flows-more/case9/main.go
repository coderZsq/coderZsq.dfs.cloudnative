package main

var t interface{}

func panic(v interface{}) {
	t = v
}

func recover() interface{} {
	return t
}
