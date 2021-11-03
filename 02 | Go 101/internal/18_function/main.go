package main

func init() {
	//func (a int, b, c string) (x, y int, z bool)
	//func (x int, y, z string) (a, b int, c bool)
	//func (_ int, _, _ string) (_, _ int, _ bool)
	//
	//func (int, string, string) (int, int, bool) // 标准函数字面形式
	//func (a int, b string, c string) (int, int, bool)
	//func (x int, _ string, z string) (int, int, bool)
	//func (int, string, string) (x int, y int, z bool)
	//func (int, string, string) (a int, b int, _ bool)
}

func init() {
	//// 这三个函数类型字面形式是等价的。
	//func () (x int)
	//func () (int)
	//func () int
	//
	//// 这两个函数类型字面形式是等价的。
	//func (a int, b string) ()
	//func (a int, b string)
}