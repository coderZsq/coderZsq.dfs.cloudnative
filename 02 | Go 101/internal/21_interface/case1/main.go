package main

// 一个非定义接口类型。
var _ interface {
	About() string
}

// ReadWriter是一个定义的接口类型。
type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

// Runnable是一个非定义接口类型的别名。
type Runnable = interface {
	Run()
}

func main() {

}