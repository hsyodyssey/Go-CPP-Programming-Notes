# Go/CPP-Programming-Notes

## Introduction

This repo records the daily programming nodes of Golang.

### Why Golang?

Golang是一种非常适合分布式系统/数据库(Distributed System/database)，区块链(Blockchain)项目开发的编程语言，因为:

- **语法简单**。它学习曲线平滑。有其他高级语言编程经验的开发人员，可以快速上手。
- **高并发基因**。它提供了强大简洁易用稳定的标准库，对高并发，分布式程序有灵活简洁全面的标准库支持。不需要依赖第三方库就可以快速构建后端/网络/高并发应用。
- **跨平台**。它支持跨系统交叉编译。MacOS/Win 也可以编译Linux下可以运行的程序，不需要担心环境依赖的问题。
- **新**。相对于C/C++，Golang的设计理念更新，历史包袱更少。对各种~~新型数据结构~~比如JSON, CSV提供了简洁流畅的原生标准库支持。
- **快**。它相比C++编译快，相比Python运行快。

所以Go非常适合做一些分布式，区块链领域的学术/工程项目的原型(Prototype)开发工作。

### 不足

- **轮子少**。相比于Java生态系统中充足的轮子，Go社区中的轮子的发展还属于初级阶段。
- **代码冗余**。相比于C++，Golang对范型(Generic Programming)支持较差。针对相同的逻辑，在面对不同的数据结构仍需要编写大量重复的代码。

## General Tips

- 保持包名和目录名的一致。
- 包名尽量简短，应该为小写单词，不要使用下划线或者驼峰式命名。
- 文件名为小写单词，使用下划线分割。
- Go语言区分大小写，小写字母开头的单词表示仅为包内访问，相当于Java/C++中的Private。跨包使用的成员变量/函数要使用大写字母开头。
- 变量和结构体的命名使用驼峰式。
- 通过更改首字母的大小写，来控制结构体的使用场景。
- 常量通常使用全部大写字母。
- Go结构体默认的小写字母开头的变量不是public的不能跨包访问。
- 通过go run *.go || go build .的方式来运行多文件同package调度的go程序
- Go语言中包括数组和Slice两种集合型数据结构。
- Go语言中的数据大小是固定的。
- Go语言中中Slice(切片)的大小可以是动态的。
- Go语言中，可以定义一个不指明长度的切片。
- 二维切片需要初始化:

```go
    //2-D Slice init
    // 初始化一个n行m列的切片
    var K[][] int.
    for i := 0; i <=n; i++ {
        inline := make([]int, m)
        K := append(K, inline)
    }
```

- Go 自带了解析CSV的库 “encoding/csv”
- Go Interface 是一种type
- 要实现Interface 必须实现interface中所有的方法
- 使用runtime Package中的 MemStats 和 ReadMemStats 来测量当前程序中Memory的使用状况。
- Context 是 Go 1.7 之后新引入的标准库接口。
- sync/atomic标准库包中提供的原子操作，通常是无锁的。
- Function Types: A function type denotes the set of all functions with the same parameter and result types.
- 当初始化结构体时，对于结构体中复杂的成员，需要显式的声明它的类型信息，否则编译无法通过。
  - For example:
  
  ```Go
  var pd1 PaperDetails = PaperDetails{
  Title:  "testtitle",
  Topics: []string{"test topic1", "test topic2"},

  Authors:      []string{"test author1", "test author2"},
  Comments:     []string{"test comment1"},
  Keysentencts: map[string][]Sentence{"test topic1": {sent1, sent2}, "test topic2": {sent2, sent3}},
  }
  ```

## Go Struct

Struct是是Go语言中最重要的自定义的数据结构之一。它的用法与C++/Java中的class既有相同点又有不同点。

- Struct 有两种初始化方式:
  - 指针的方式, 这种方式会返回一个指向新的结构体的指针:

    ``` Golang
    var s * Student
    s = new(Student)
    s.name = "test"

    \\or
    var s = new (Student)
    s.name = "test"
    ```

  - With a struct literal, 这种方式会返回一个值变量(等于第一种方式的*s):

    ```Golang
    s := Student{
      name: "test",
    }
    ```

## Map

- 可以通过第二个返回值的方式来判断map中是否存在目标key。
- Go 中的Map是不能被Copy的，实现Copy的方法是循环赋值。
- Map在调用的时候可能会存在浅拷贝的问题。浅拷贝的意思是指，新对象被赋予了与旧对象相同的指针地址，当旧对象的值发生变化的时候，新对象的值也会发生变化。
- 在使用Map和Slice的时候，要注意浅拷贝带来的赋值错误的问题。浅拷贝发生时，赋值对象的双方指向的是同一个内存地址。
- map的值可以是一个函数。关于这个特征可以参考的我们的[实例代码](./example/map/case_function/main.go)

```go
m := make(map[string]sting)
m["001"] = "test001"
m["002"] = "test002"
value, isExist := m["001"] // value: string: test001, isExist: bool : true
value, isExist := m["003"] // value: , isExist: bool : false

```

- Go 语言原生的map类型不支持并发的读写操作。一个第三方的解决方案, [concurrent-map](https://github.com/orcaman/concurrent-map).

## 内存管理

- Go的堆(Heap)是从0xc000000000(40bits，10个hex)位置开始增长的。相关的讨论:<https://github.com/golang/go/issues/27583>
  - 关于Address space layout randomization(ASLR)的介绍：<https://en.wikipedia.org/wiki/Address_space_layout_randomization>

## 并发(goroutine)

- Go原生支持协程(Goroutine)，使得go语言天生适合高并发的应用。
- 并发避免不了的会潜在Data Race的问题，可以通过go run -race xx .go/ go build -race xx .go来检测程序中的data race 问题。
- 对于可能产生Data race的代码，尽量使用atomic类型，或者加锁Mutex.

## Go Channel

- 在go语言中，Channel主要用于goroutine之间的通信。
- 同时，Channel还可以起到阻塞的作用。例如，我们可以声明一个Bool类型的Channel: c := make(chan bool)。如果我们将`<-c`放在函数A中，那么在`c`接受到之前，函数A将会一直处于阻塞(Block)的状态.
- `chan<- int`表示send-only Channel，只能接收数据。
- `<-chan int`表示receive-only的 Channel，只能向外传输数据。
- Channel其实是一个先入先出(FIFO)的队列。

## Go File

- Go file.ReadAt的实现中使用了一个for循环来反复的向目标数组b中传入数据。不知道为什么要这么设计，而不是一次性的从文件中拿出所需要的数据。

  ```go
   for len(b) > 0 {
  m, e := f.pread(b, off)
  if e != nil {
   err = f.wrapErr("read", e)
   break
  }
  n += m
  b = b[m:]
  off += int64(m)

 }
  ```

## FQA

- What is struct{} and struct{}{} ?
  - struct{} 表示一个零元素的struct结构。通常会被用在没有任何信息被存储的场景中。
  - struct{}{} 表示一个存储了struct{}的Composite literal.
  - Source: [Stackoverflow](https://stackoverflow.com/questions/45122905/how-do-struct-and-struct-work-in-go)
  - interface是一种万能的数据类型，他可以接受任何类型的值。

## Go Unit Testing

- Go 单元测试中，执行到t.Error()/ t.Errorf() 测试函数会输出错误的log信息并继续执行。
- Go 单元测试中，执行到t.Fatal()/ t.Fatalf() 测试函数会输出错误的log信息并结束测试。
- Go 单元测试中，如果不满足断言assert.Equal() 测试函数会结束并报错。

## Go 性能分析

- PProf
- Trace
  - go tool trace trace.out

## Go and C and Python

- Python 可以调用Go编译的动态链接库 and.so (Shared Object)
- Go编译成so文件时，需要在函数上一行添加//export xxx(函数名).
  - 注意//与export中间不能有空格
  - Example 位于example/pygo

## Tricks

- 利用CPU Cache的特性来加速程序的运行。 目前x86 CPU Cache line的长度为64 bytes，CPU Cache缓存数据通常以Cache line为单位。可以使用Cache Padding的方式，通过额外添加数据结构，构造64 bytes的结构来降低CPU Cache中False Sharing的产生，提高程序运行速度。

## Log

- Log.Fatal会直接退出程序，不会执行defer相关的函数

## Reference

- Go语言高性能编程 [link](https://geektutu.com/post/high-performance-go.html)
- Go By Example [link](https://gobyexample.com/)
- <https://www.pengrl.com/p/9125/>
- Go PProf [[link]](https://segmentfault.com/a/1190000016412013)
- Go Trace [[link]](https://segmentfault.com/a/1190000019736288)
