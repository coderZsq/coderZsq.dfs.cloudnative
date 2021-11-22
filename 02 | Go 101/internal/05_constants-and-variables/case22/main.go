package main

var x, y, z = 123, true, "foo" // 包级变量

func main() {
	var q, r = 789, false
	r, s := true, "bar"
	r = y // r没有被有效使用。
	x = q // q被有效使用了。

	_, _ = r, s // 将r和s做为源值使用一次。
}
