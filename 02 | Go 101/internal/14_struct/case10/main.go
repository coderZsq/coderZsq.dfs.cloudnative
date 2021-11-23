package main

type S0 struct {
	y int "foo"
	x bool
}

type S1 = struct { // S1是一个非定义类型
	x int "foo"
	y bool
}

type S2 = struct { // S2也是一个非定义类型
	x int "bar"
	y bool
}

type S3 S2 // S3是一个定义类型。
type S4 S3 // S4是一个定义类型。
// 如果不考虑字段标签，S3（S4）和S1的底层类型一样。
// 如果考虑字段标签，S3（S4）和S1的底层类型不一样。

var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}
func f() {
	v1 = S1(v2); v2 = S2(v1)
	v1 = S1(v3); v3 = S3(v1)
	v1 = S1(v4); v4 = S4(v1)
	v2 = v3; v3 = v2 // 这两个转换可以是隐式的
	v2 = v4; v4 = v2 // 这两个转换也可以是隐式的
	v3 = S3(v4); v4 = S4(v3)
}

func main() {

}
