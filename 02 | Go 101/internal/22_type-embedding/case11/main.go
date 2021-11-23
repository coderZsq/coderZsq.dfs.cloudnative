package main

// 注意：这些声明不是合法的Go语法。这里这样表示只是为了
// 解释目的。它们有助于解释提升方法值是如何被估值的。
func (s Singer) PrintName = s.Person.PrintName
func (s *Singer) PrintName = s.Person.PrintName
func (s *Singer) SetAge = s.Person.SetAge

func main() {

}
