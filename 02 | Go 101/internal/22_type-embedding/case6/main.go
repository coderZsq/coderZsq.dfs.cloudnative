package main


type A struct {
	FieldX int
}

func (a A) MethodA() {}

type B struct {
	*A
}

type C struct {
	B
}

func main() {
	var c = &C{B: B{A: &A{FieldX: 5}}}

	// 这几行是等价的。
	_ = c.B.A.FieldX
	_ = c.B.FieldX
	_ = c.A.FieldX // A是类型C的一个提升字段
	_ = c.FieldX   // FieldX也是一个提升字段

	// 这几行是等价的。
	c.B.A.MethodA()
	c.B.MethodA()
	c.A.MethodA()
	c.MethodA() // MethodA是类型C的一个提升方法
}