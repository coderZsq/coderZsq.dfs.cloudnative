package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//if InitSimpleStatement; Condition {
	//	// do something
	//} else {
	//	// do something
	//}
}

func init() {
	rand.Seed(time.Now().UnixNano())

	if n := rand.Int(); n%2 == 0 {
		fmt.Println(n, "是一个偶数。")
	} else {
		fmt.Println(n, "是一个奇数。")
	}

	n := rand.Int() % 2 // 此n不是上面声明的n
	if n % 2 == 0 {
		fmt.Println("一个偶数。")
	}

	if ; n % 2 != 0 {
		fmt.Println("一个奇数。")
	}
}

func init() {
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("现在为上午。")
	} else if h > 19 {
		fmt.Println("现在为晚上。")
	} else {
		fmt.Println("现在为下午。")
		// 左h是一个新声明的变量，右h已经在上面声明了。
		h := h
		// 刚声明的h遮掩了上面声明的h。
		_ = h
	}

	// 上面声明的两个h在此处都不可见。
}

func init() {
	//for InitSimpleStatement; Condition; PostSimpleStatement {
	//	// do something
	//}
}

func init() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func init() {
	var i = 0
	for ; i < 10; {
		fmt.Println(i)
		i++
	}
	for i < 20 {
		fmt.Println(i)
		i++
	}
}

func init() {
	for i := 0; ; i++ { // 等价于：for i := 0; true; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}

	// 下面这几个循环是等价的。
	//for ; true; {
	//}
	//for true {
	//}
	//for ; ; {
	//}
	//for {
	//}
}

func init() {
	for i := 0; i < 3; i++ {
		fmt.Print(i)
		i := i // 这里声明的变量i遮挡了上面声明的i。
		// 右边的i为上面声明的循环变量i。
		i = 10 // 新声明的i被更改了。
		_ = i
	}
}

func init() {
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func init() {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Print(i)
	}
}

func init() {
	//switch InitSimpleStatement; CompareOperand0 {
	//case CompareOperandList1:
	//	// do something
	//case CompareOperandList2:
	//	// do something
	//	...
	//case CompareOperandListN:
	//	// do something
	//default:
	//	// do something
	//}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n%9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")

		// 和很多其它语言不一样，程序不会自动从一个
		// 分支代码块跳到下一个分支代码块去执行。
		// 所以，这里不需要一个break语句。
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
		break // 这里的break语句可有可无的，效果
		// 是一样的。执行不会跳到下一个分支。
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	// case 6, 7, 8:
	// 上一行可能编译不过，因为6和上一个case中的
	// 6重复了。是否能编译通过取决于具体编译器实现。
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		fallthrough // 跳到下个代码块
	case 5, 6, 7, 8:
		// 一个新声明的n，它只在当前分支代码快内可见。
		n := 99
		fmt.Println("n =", n) // 99
		fallthrough
	default:
		// 下一行中的n和第一个分支中的n是同一个变量。
		// 它们均为switch表达式"n"。
		fmt.Println("n =", n)
	}
}

func init() {
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		// 此整个if代码块为当前分支中的最后一条语句
		if true {
			//fallthrough // error: 不是当前分支中的最后一条语句
		}
	case 5, 6, 7, 8:
		n := 99
		//fallthrough // error: 不是当前分支中的最后一条语句
		_ = n
	default:
		fmt.Println(n)
		//fallthrough // error: 不能出现在最后一个分支中
	}
}

func main() {}