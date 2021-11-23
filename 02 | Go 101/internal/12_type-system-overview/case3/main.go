package main

// 下面这些新定义的类型和它们的源类型都是基本类型。
type (
	MyInt int
	Age   int
	Text  string
)

// 下面这些新定义的类型和它们的源类型都是组合类型。
type IntPtr *int
type Book struct{author, title string; pages int}
type Convert func(in0 int, in1 bool)(out0 int, out1 string)
type StringArray [5]string
type StringSlice []string

func f() {
	// 这三个新定义的类型名称只能在此函数内使用。
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{Read([]byte) int}
}

func main() {

}