package main

import (
	"fmt"
	"sync"
)

type Fruit struct {
	num  int
	lock sync.Mutex // Struct Lock
}

func main() {
	var apple Fruit

	apple.num = 0

	for {
		// Create Apple
		go func() {
			apple.lock.Lock()
			if apple.num == 0 {
				apple.num++
				fmt.Println("Create Apple: ", apple.num)
			}
			apple.lock.Unlock()

		}()

		// Consume Apple
		go func() {
			apple.lock.Lock()
			if apple.num == 1 {
				apple.num--
				fmt.Println("Eat Apple: ", apple.num)
			}
			apple.lock.Unlock()
		}()

	}

}
