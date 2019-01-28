package main

import "fmt"

type books struct {
	title string
	author string
	price float64
}

func main() {
	var b books
	b.author="zhangsan"
	b.title="标题"
	b.price=100
	fmt.Println(b.title)
	printBooks(b)
}
func printBooks(books books)  {
	fmt.Printf(books.title)
	fmt.Printf(books.author)
	fmt.Println(books.price)

}
