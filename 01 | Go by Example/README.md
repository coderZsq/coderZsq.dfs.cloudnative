## Go by Example

- [Go by Example](https://gobyexample.com/)
- [Go by Example 中文版 (gobyexample-cn.github.io)](https://gobyexample-cn.github.io/)

## Reference

- [Go Slices: usage and internals - go.dev](https://go.dev/blog/slices-intro)
- [Error handling and Go - go.dev](https://go.dev/blog/error-handling-and-go)
- [how-to-use-interfaces-in-go - jordan orelli](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
- [research!rsc: Go Data Structures (swtch.com)](https://research.swtch.com/godata)
- [research!rsc: Go Data Structures: Interfaces (swtch.com)](https://research.swtch.com/interfaces)
- [Frequently Asked Questions (FAQ) - The Go Programming Language (golang.org)](https://golang.org/doc/faq#closures_and_goroutines)
- [Strings, bytes, runes and characters in Go - go.dev](https://go.dev/blog/strings)
- [JSON and Go - go.dev](https://go.dev/blog/json)

## Go Data Structures - research!rsc

### Basic types

让我们从一些简单的例子开始。

![godata1.png (409×303) (swtch.com)](https://research.swtch.com/godata1.png)

变量i的类型为int，在内存中表示为一个32位的字。(所有这些图片显示的是32位的内存布局；在目前的实现中，只有指针在64位机器上会变大--int仍然是32位--尽管一个实现可以选择使用64位来代替。)

变量j的类型是int32，因为有显式转换。尽管i和j有相同的内存布局，但它们有不同的类型：赋值i = j是一个类型错误，必须用显式转换来写：i = int（j）。

变量f的类型是float，目前的实现将其表示为32位浮点值。它的内存占用与int32相同，但内部布局不同。

### Structs and pointers

现在事情开始有了起色。变量bytes的类型为[5]byte，是一个5字节的数组。它的内存表示就是这5个字节，一个接一个，就像C语言的数组。同样地，primes是一个4个字节的数组。

Go与C语言一样，但与Java不同，它让程序员控制什么是指针，什么不是指针。例如，这个类型定义。

```go
type Point struct {X, Y int}
```

定义了一个名为Point的简单结构类型，在内存中表示为两个相邻的ints。

![godata1a.png (285×151) (swtch.com)](https://research.swtch.com/godata1a.png)

复合字面语法Point{10, 20}表示一个初始化的Point。以复合字面的地址表示一个指向新分配和初始化的Point的指针。前者是内存中的两个字；后者是一个指向内存中两个字的指针。

结构中的字段在内存中是并排排列的。

```go
type Rect1 struct {Min, Max Point}
type Rect2 struct {Min, Max *Point}
```

![godata1b.png (443×151) (swtch.com)](https://research.swtch.com/godata1b.png)

Rect1是一个有两个Point字段的结构，由两个Powers--四个ints--组成的一行表示。Rect2是一个有两个*Point字段的结构，由两个*Points表示。

使用过C语言的程序员可能不会对点字段和*点字段之间的区别感到惊讶，而只使用过Java或Python（或......）的程序员可能会对必须做出这个决定感到惊讶。通过让程序员控制基本的内存布局，Go提供了控制特定数据结构集合的总大小、分配数量和内存访问模式的能力，所有这些对于构建性能良好的系统都很重要。

### Strings

有了这些初步的认识，我们就可以转向更有趣的数据类型。

![godata2.png (264×151) (swtch.com)](https://research.swtch.com/godata2.png)

(灰色的箭头表示在实现中存在的指针，但在程序中不直接可见。)

一个字符串在内存中表示为一个2字结构，包含一个指向字符串数据的指针和一个长度。因为字符串是不可改变的，所以多个字符串共享同一个存储空间是安全的，所以切片的结果是一个新的2字结构，其指针和长度可能不同，但仍然指向同一个字节序列。这意味着切分可以在不分配或复制的情况下进行，使字符串切分和传递显式索引一样高效。

(另外，在Java和其他语言中，有一个众所周知的问题，当你把一个字符串切成小块保存时，对原始字符串的引用会把整个原始字符串保留在内存中，尽管只有一小部分仍然需要。Go也有这个问题。我们曾尝试过另一种方法，但被拒绝了，那就是让字符串切分变得如此昂贵--一次分配和一次复制--以至于大多数程序都会避免这样做）。)

### Slices

![godata3.png (470×151) (swtch.com)](https://research.swtch.com/godata3.png)

一个片断是对一个数组的一个部分的引用。在内存中，它是一个3个字的结构，包括第一个元素的指针，分片的长度和容量。长度是像x[i]这样的索引操作的上限，而容量是像x[i:j]这样的分片操作的上限。

和切分字符串一样，切分数组并不产生副本：它只是创建一个新的结构，持有不同的指针、长度和容量。在这个例子中，评估复合字面[]int{2, 3, 5, 7, 11}会创建一个包含五个值的新数组，然后设置分片x的字段来描述该数组。分片表达式x[1:3]并没有分配更多的数据：它只是写了一个新的分片结构的字段来引用相同的后备存储。在这个例子中，长度是2-y[0]和y[1]是唯一有效的索引-但是容量是4-y[0:4]是一个有效的分片表达式。(关于长度和容量以及如何使用分片的更多信息，请参见Effective Go。)

因为分片是多字结构，而不是指针，所以分片操作不需要分配内存，甚至不需要分配分片头，分片头通常可以保留在堆栈中。这种表示方法使得切片的使用与C语言中传递显式指针和长度对一样便宜。Go最初将切片表示为指向上图所示结构的指针，但这样做意味着每个切片操作都要分配一个新的内存对象。即使有一个快速的分配器，这也会给垃圾收集器带来很多不必要的工作，我们发现，就像上面字符串的情况一样，程序会避免切分操作而选择传递显式索引。去掉中介和分配，使得切片的成本足够低，在大多数情况下可以避免传递显式索引。

### New and Make

Go有两个数据结构创建函数：new和make。这种区别是早期常见的混淆点，但似乎很快就变得自然了。基本的区别是new(T)返回一个*T，一个Go程序可以隐式解除引用的指针（图中的黑色指针），而make(T, args)返回一个普通的T，不是一个指针。通常这个T里面有一些隐式指针（图中的灰色指针）。New返回一个指向归零内存的指针，而make返回一个复杂的结构。

![godata4.png (470×627) (swtch.com)](https://research.swtch.com/godata4.png)

有一种方法可以将这两者统一起来，但这将是对C和C++传统的重大突破：定义make(*T)来返回一个指向新分配的T的指针，这样，当前的new(Point)将被写成make(*Point)。我们尝试了几天，但认为这与人们对分配函数的期望差别太大。

### Interfaces

#### Interface Values

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

#### Computing the Itable

现在我们知道 itables 是什么样子的了，但是它们是从哪里来的呢？Go 的动态类型转换意味着编译器或链接器预先计算所有可能的 itables 是不合理的：有太多的（接口类型，具体类型）对，而且大多数都不需要。相反，编译器为每个具体类型生成一个类型描述结构，比如二进制或 int 或 func(map[int]string)。在其他元数据中，类型描述结构包含一个由该类型实现的方法列表。同样地，编译器为每个接口类型（如 Stringer）生成一个（不同的）类型描述结构；它也包含一个方法列表。接口运行时通过在具体类型的方法表中寻找接口类型的方法表所列的每个方法来计算 itable。运行时在生成 itable 后对其进行缓存，因此这种对应关系只需要计算一次。

在我们的简单例子中，Stringer 的方法表有一个方法，而 Binary 的方法表有两个方法。一般来说，接口类型可能有 ni 个方法，具体类型有 nt 个方法。显然，寻找从接口方法到具体方法的映射需要 O(ni × nt)时间，但我们可以做得更好。通过对这两个方法表进行排序并同时走动，我们可以在 O(ni + nt)时间内建立映射关系。

#### Memory Optimizations

上述实施所使用的空间可以通过两种互补的方式进行优化。

首先，如果所涉及的接口类型是空的--它没有任何方法--那么 itable 除了持有指向原始类型的指针外没有任何作用。在这种情况下，itable 可以被丢弃，值可以直接指向该类型。

![](https://research.swtch.com/gointer3.png)

一个接口类型是否有方法是一个静态属性--要么源代码中的类型写着 interface{}，要么写着 interace{ methods... }，所以编译器知道在程序的每一点上使用的是哪种表示法。

其次，如果与接口值相关的值可以装在一个机器字中，就没有必要引入中介或堆分配。如果我们将 Binary32 定义为像 Binary 一样，但实现为 uint32，就可以通过将实际值保留在第二个字中来存储在一个接口值中。

![](https://research.swtch.com/gointer4.png)

实际值是被指向还是被内联，取决于类型的大小。编译器安排在类型的方法表中列出的函数（这些函数被复制到 itables 中）对被传入的字做正确的事情。如果接收者类型适合一个字，它就被直接使用；如果不适合，它就被取消引用。图中显示了这一点：在上面的二进制版本中，itable 中的方法是(*Binary).String，而在 Binary32 的例子中，itable 中的方法是 Binary32.String 而不是(*Binary32).String。

当然，持有字大小（或更小）数值的空接口可以利用这两种优化。

![](https://research.swtch.com/gointer5.png)

#### Method Lookup Performance

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
