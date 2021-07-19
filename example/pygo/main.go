package main

import "C"
import "fmt"

//export PrintDll
func PrintDll() {
	fmt.Println("Hi Python, this message is from Go")
}

//export Sum
func Sum(a int, b int) int {
	return a + b
}

func main() {

}
