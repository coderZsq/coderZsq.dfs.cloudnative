package main

func StatRandomNumbers(numRands int) (int, int)
{ // 编译错误：语法错误
	var a, b int
	for i := 0; i < numRands; i++
	{ // 编译错误：语法错误
		if rand.Intn(MaxRand) < MaxRand/2
		{ // 编译错误：语法错误
			a = a + 1
		} else {
			b++
		}
	}
	return a, b
}
