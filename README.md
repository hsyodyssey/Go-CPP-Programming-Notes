# Go-Developer-Tips

Daily Programming Tips in Go

## General Tips

- 保持包名和目录名的一致。
- 包名尽量简短，应该为小写单词，不要使用下划线或者驼峰式命名。
- 文件名为小写单词，使用下划线分割。
- Go语言区分大小写，小写字母开头的单词为包内访问，相当于Java/C++中的Private。跨包使用的成员变量/函数要使用大写字母开头。
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

```2-D Slice init
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

## Go and C and Python

- Python 可以调用Go编译的动态链接库 and.so (Shared Object) 
- Go编译成so文件时，需要在函数上一行添加//export xxx(函数名).
  - 注意//与export中间不能有空格
  - Example 位于example/pygo

## Log

- Log.Fatal会直接退出程序，不会执行defer相关的函数




