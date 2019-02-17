package main

import "fmt"

type A struct {
	name string
}
type B struct {
	name string
}

func main() {
	a := A{}
	a.name = "aa2"
	a.print()
	fmt.Println(a.name)
	b := B{}
	b.name = "bb2"
	b.print()
	fmt.Println(b.name)

}
func (a *A) print() {
	a.name = "aa"
	fmt.Println("aa1")

}
func (b B) print() {
	b.name = "bb"
	fmt.Println("bb1")

}
