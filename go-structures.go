package main

import "fmt"

type books struct {
	title string
	author string
	price float64
}
type books1 struct {
	//结构体嵌套，匿名成员
	books
	brief string
}

func main() {
	var b books
	b.author="zhangsan"
	b.title="标题"
	b.price=100
	fmt.Println(b.title)
	printBooks(b)
	var b1 books1
	b1.brief="描述"
	fmt.Println(b1.brief)
}
func printBooks(books books)  {
	fmt.Printf(books.title)
	fmt.Printf(books.author)
	fmt.Println(books.price)

}
