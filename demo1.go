package main

import "fmt"

func main() {
	//var g int =9
	var a  =1;
	var b  = 2;
	var c = 0;
	fmt.Println("main函数中a的值为：",a)
	c = sum(a,b)
	fmt.Println("main函数中b的值为：",b)
	fmt.Println("main函数中c的值为：",c)
}
func sum(a,b int) int {
	fmt.Println("sum()函数中a的值为：",a)
	fmt.Println("sum()函数中b的值为：",b)
	return a+b

}


