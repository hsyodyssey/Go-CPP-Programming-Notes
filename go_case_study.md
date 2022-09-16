# Golang 相关的实例

## 浅拷贝/深拷贝问题(Deep and Shallow Copy)

## Go 协程实例

### Case 1

实现一个这样的功能：有限时间赛马:

- 主线程设置一个**约定时间**，并发多个goroutine去执行任务，
- 当**约定时间**超时时，结束所有的goroutine并返回。
- 当goroutine执行完之后直接返回。

具体的代码可以在[这里](example/goroutine/case_1/main.go)

### Case 2 Go程序的逃逸分析

具体的代码例子请看[这里](example/compiler/escape_heap/main.go)

Q: `go build main.go`的之后的运行结果是什么? `go build -gcflags="-N" main.go`之后的运行结果是什么?

A: true false and true true

原代码中存在堆逃逸问题，该部分的代码`f1 == f2`会被编译器直接优化成false。

## Map的一些问题

- Go 中的Map是不能被Copy的，实现Copy的方法是循环赋值。
