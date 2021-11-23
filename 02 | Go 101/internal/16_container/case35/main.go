package main

func main() {
	// 忽略键值循环变量。
	for _, element = range aContainer {
		// ...
	}

	// 忽略元素循环变量。
	for key, _ = range aContainer {
		element = aContainer[key]
		// ...
	}

	// 舍弃元素循环变量。此形式和上一个变种等价。
	for key = range aContainer {
		element = aContainer[key]
		// ...
	}

	// 键值和元素循环变量均被忽略。
	for _, _ = range aContainer {
		// 这个变种形式没有太大实用价值。
	}

	// 键值和元素循环变量均被舍弃。此形式和上一个变种等价。
	for range aContainer {
		// 这个变种形式没有太大实用价值。
	}
}