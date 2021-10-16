## Go by Example

- [Go by Example](https://gobyexample.com/)

## Reference

- [Go Slices: usage and internals - go.dev](https://go.dev/blog/slices-intro)
- [Error handling and Go - go.dev](https://go.dev/blog/error-handling-and-go)
- [how-to-use-interfaces-in-go - jordan orelli](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
- [research!rsc: Go Data Structures: Interfaces (swtch.com)](https://research.swtch.com/interfaces)
- [Frequently Asked Questions (FAQ) - The Go Programming Language (golang.org)](https://golang.org/doc/faq#closures_and_goroutines)

### Interface Values

有方法的语言通常分为两个阵营：为所有的方法调用准备静态表格（如 C++和 Java），或者在每次调用时进行方法查找（如 Smalltalk 及其许多模仿者，包括 JavaScript 和 Python），并添加花哨的缓存以使该调用有效。Go 介于这两者之间：它有方法表，但在运行时计算它们。我不知道 Go 是否是第一个使用这种技术的语言，但它肯定不是一种常见的技术。

作为热身，二进制类型的值只是一个由两个 32 位字组成的 64 位整数（和上一篇文章一样，我们将假设一台 32 位机器；这次内存是向下增长而不是向右增长）。

![](https://research.swtch.com/gointer1.png)

接口值被表示为两个字对，给出一个指向存储在接口中的类型信息的指针和一个指向相关数据的指针。将 b 分配给一个类型为 Stringer 的接口值，可以设置接口值的两个字。

![](https://research.swtch.com/gointer2.png)

(接口值中包含的指针是灰色的，以强调它们是隐含的，没有直接暴露给 Go 程序)。

接口值中的第一个词指向我所说的接口表或 itable（发音为 i-table；在运行时源中，C 语言的实现名称为 Itab）。itable 以一些关于相关类型的元数据开始，然后变成一个函数指针的列表。注意，itable 对应的是接口类型，而不是动态类型。就我们的例子而言，持有 Binary 类型的 Stringer 的 itable 列出了用于满足 Stringer 的方法，它只是 String。Binary 的其他方法（Get）在 itable 中没有出现。

接口值中的第二个字指向实际的数据，在本例中是 b 的副本。赋值 var s Stringer = b 做了一个 b 的副本，而不是指向 b，原因和 var c uint64 = b 做了一个副本一样：如果 b 后来改变了，s 和 c 应该是原来的值，而不是新的。存储在接口中的值可能是任意大的，但在接口结构中只有一个字是专门用来保存该值的，所以赋值在堆上分配了一大块内存，并将指针记录在一个字槽中。(当值确实适合在槽中时，有一个明显的优化；我们将在后面讨论这个问题)。

为了检查一个接口值是否持有一个特定的类型，就像上面的类型切换一样，Go 编译器会生成相当于 C 语言表达式 s.tab->type 的代码，以获得类型指针，并将其与所需的类型进行检查。如果类型匹配，就可以通过解引用 s.data 来复制该值。

为了调用 s.String()，Go 编译器生成的代码相当于 C 语言表达式 s.tab->fun\[0](s.data)：它从 itable 中调用适当的函数指针，将接口值的数据字作为函数的第一个（在这个例子中，唯一的）参数传递给它。如果你运行 8g -S x.go，你可以看到这段代码（细节在这篇文章的底部）。注意，itable 中的函数被传递了来自接口值第二个字的 32 位指针，而不是它指向的 64 位值。一般来说，接口调用站点不知道这个词的含义，也不知道它指向的数据有多少。相反，接口代码安排，itable 中的函数指针期待着存储在接口值中的 32 位表示。因此，本例中的函数指针是（\*Binary）.String 而不是 Binary.String。

我们正在考虑的例子是一个只有一个方法的接口。一个有更多方法的接口在 itable 底部的 fun 列表中会有更多条目。

### Computing the Itable

现在我们知道 itables 是什么样子的了，但是它们是从哪里来的呢？Go 的动态类型转换意味着编译器或链接器预先计算所有可能的 itables 是不合理的：有太多的（接口类型，具体类型）对，而且大多数都不需要。相反，编译器为每个具体类型生成一个类型描述结构，比如二进制或 int 或 func(map[int]string)。在其他元数据中，类型描述结构包含一个由该类型实现的方法列表。同样地，编译器为每个接口类型（如 Stringer）生成一个（不同的）类型描述结构；它也包含一个方法列表。接口运行时通过在具体类型的方法表中寻找接口类型的方法表所列的每个方法来计算 itable。运行时在生成 itable 后对其进行缓存，因此这种对应关系只需要计算一次。

在我们的简单例子中，Stringer 的方法表有一个方法，而 Binary 的方法表有两个方法。一般来说，接口类型可能有 ni 个方法，具体类型有 nt 个方法。显然，寻找从接口方法到具体方法的映射需要 O(ni × nt)时间，但我们可以做得更好。通过对这两个方法表进行排序并同时走动，我们可以在 O(ni + nt)时间内建立映射关系。

### Memory Optimizations

上述实施所使用的空间可以通过两种互补的方式进行优化。

首先，如果所涉及的接口类型是空的--它没有任何方法--那么 itable 除了持有指向原始类型的指针外没有任何作用。在这种情况下，itable 可以被丢弃，值可以直接指向该类型。

![](https://research.swtch.com/gointer3.png)

一个接口类型是否有方法是一个静态属性--要么源代码中的类型写着 interface{}，要么写着 interace{ methods... }，所以编译器知道在程序的每一点上使用的是哪种表示法。

其次，如果与接口值相关的值可以装在一个机器字中，就没有必要引入中介或堆分配。如果我们将 Binary32 定义为像 Binary 一样，但实现为 uint32，就可以通过将实际值保留在第二个字中来存储在一个接口值中。

![](https://research.swtch.com/gointer4.png)

实际值是被指向还是被内联，取决于类型的大小。编译器安排在类型的方法表中列出的函数（这些函数被复制到 itables 中）对被传入的字做正确的事情。如果接收者类型适合一个字，它就被直接使用；如果不适合，它就被取消引用。图中显示了这一点：在上面的二进制版本中，itable 中的方法是(*Binary).String，而在 Binary32 的例子中，itable 中的方法是 Binary32.String 而不是(*Binary32).String。

当然，持有字大小（或更小）数值的空接口可以利用这两种优化。

![](https://research.swtch.com/gointer5.png)

### Method Lookup Performance

Smalltalk 和许多紧随其后的动态系统在每次方法被调用时都会进行方法查找。为了提高速度，许多实现在每个调用点使用一个简单的单入口缓存，通常是在指令流本身。在一个多线程程序中，这些缓存必须被小心管理，因为多个线程可能同时在同一个调用点。即使避免了竞赛，缓存最终也会成为内存争夺的一个来源。

因为 Go 有静态类型的提示，配合动态方法的查找，它可以将查找从调用站点移回到值存储在接口中的位置。例如，考虑这个代码片断。

```go
var any interface{} // initialized elsewhere
s := any.(Stringer) // dynamic conversion
for i := 0; i < 100; i++ {
	fmt.Println(s.String())
}
```

在 Go 中，itable 在第 2 行的赋值过程中被计算出来（或在缓存中找到）；在第 4 行执行的 s.String()调用的调度是几个内存获取和一个间接调用指令。

相比之下，用 Smalltalk（或 JavaScript，或 Python，或......）这样的动态语言来实现这个程序会在第 4 行进行方法查找，这在循环中会重复不必要的工作。前面提到的缓存使其成本降低，但仍比单一的间接调用指令要贵。

```go
b := Binary(200)
s := Stringer(b)
fmt.Println(s.String())
```

```s
0045 (x.go:25) LEAL    s+-24(SP),BX
0046 (x.go:25) MOVL    4(BX),BP
0047 (x.go:25) MOVL    BP,(SP)
0048 (x.go:25) MOVL    (BX),BX
0049 (x.go:25) MOVL    20(BX),BX
0050 (x.go:25) CALL    ,BX
```

LEAL 将 s 的地址加载到寄存器 BX 中。（符号 n(SP)描述内存中 SP+n 的字。0(SP)可以简称为(SP))。接下来的两条 MOVL 指令从接口的第二个字中获取数值，并将其作为第一个函数调用参数 0(SP)存储。最后两条 MOVL 指令从 itable 中获取数值，然后从 itable 中获取函数指针，为调用该函数做准备。
