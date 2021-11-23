package main

func main() {
	var _ func (a int, b string, c string) (x int, y int, z bool)

	var _ func (a int, b, c string) (x, y int, z bool)

	var _ func (x int, y, z string) (a, b int, c bool)

	var _ func (_ int, _, _ string) (_, _ int, _ bool)
}
