package main

import "fmt"

func main() {

	var a *int
	var b = 20
	a = &b
	fmt.Printf("b 变量的地址是: %x\n", &b)
	fmt.Printf("a 变量的地址是: %x\n", a)
	fmt.Printf("a 变量的地址是: %x\n", *a)
}
