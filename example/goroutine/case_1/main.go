package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

/*
1. 只能编辑 foo 函数
2. foo 必须要调用 slow 函数
3. foo 函数在 ctx 超时后必须立刻返回
4. 【加分项】如果 slow 结束的比 ctx 快，也立刻返回
*/

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = trace.Start(f)

	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	foo(ctx)
	fmt.Println("Done")

}

func foo(ctx context.Context) {
	ch := make(chan struct{}, 1)

	go func() {
		slow()
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Main Context Time out")
		return
	case <-ch:
		fmt.Println("The Slow function Done")
		return
	}

}

func slow() {
	n := rand.Intn(3)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("sleep %ds\n", n)
}
