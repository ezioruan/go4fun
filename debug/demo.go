package main

import (
	"fmt"
	"runtime"
)

func test() {
	fmt.Print("aaa")
	runtime.Breakpoint()
	fmt.Print("bbb")
	fmt.Print("bb")
	fmt.Print("bb2")
	fmt.Print("bbb")
}

func main() {
	test()
}
