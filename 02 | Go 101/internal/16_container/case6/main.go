package main

// heads为一个切片值。它的类型的元素类型为*[4]byte。
// 此元素类型为一个基类型为[4]byte的指针类型。
// 此指针基类型为一个元素类型为byte的数组类型。
var heads = []*[4]byte{
	&[4]byte{'P', 'N', 'G', ' '},
	&[4]byte{'G', 'I', 'F', ' '},
	&[4]byte{'J', 'P', 'E', 'G'},
}

func main() {

}
