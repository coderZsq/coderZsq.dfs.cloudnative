package main

func main() {
	baseContainer[low : high]       // 双下标形式
	baseContainer[low : high : max] // 三下标形式

	baseContainer[low : high : cap(baseContainer)]

	// 双下标形式
	0 <= low <= high <= cap(baseContainer)

	// 三下标形式
	0 <= low <= high <= max <= cap(baseContainer)

	baseContainer[0 : len(baseContainer)]
	baseContainer[: len(baseContainer)]
	baseContainer[0 :]
	baseContainer[:]
	baseContainer[0 : len(baseContainer) : cap(baseContainer)]
	baseContainer[: len(baseContainer) : cap(baseContainer)]
}