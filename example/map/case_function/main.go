package main

import "fmt"

type mapTestfunc func(s string)

var funcMapTest = map[string]mapTestfunc{
	"one": mapTestOne,
	"two": mapTestTwo,
}

func main() {
	for k, f := range funcMapTest {
		f(k)
	}
}

func mapTestOne(s string) {
	fmt.Println("Test from Function one", s)
}

func mapTestTwo(s string) {
	fmt.Println("Test from Function two", s)
}
