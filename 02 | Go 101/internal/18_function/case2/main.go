package main

func main() {
	var _ func (int, string, string) (int, int, bool) // 标准函数字面形式
	var _ func (a int, b string, c string) (int, int, bool)
	var _ func (x int, _ string, z string) (int, int, bool)
	var _ func (int, string, string) (x int, y int, z bool)
	var _ func (int, string, string) (a int, b int, _ bool)
}