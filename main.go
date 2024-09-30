package main

import (
	"fmt"

	go_hello_world "github.com/torikallcode/module-hello-world/v2"
)

func main() {
	hello := go_hello_world.HelloWorld("Akbar")
	fmt.Println(hello)
}
