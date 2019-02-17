package main

import "fmt"

type a struct {
	name string
}
type b struct {
	name string
}

func main() {
	a:=a{}
	a.name="aa2"
	a.print()
	fmt.Println(a.name)
	b:=b{}
	b.name="bb2"
	b.print()
	fmt.Println(b.name)

}
func (a *a)print()  {
	a.name="aa"
	fmt.Println("aa1")

}
func (b b)print()  {
	b.name="bb"
	fmt.Println("bb1")

}