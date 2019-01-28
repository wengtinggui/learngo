package main

import "fmt"



func main() {
	a:=[]int{1,2,3,4,5}
	fmt.Println(a)
	fmt.Println(a[1:3])
	fmt.Println(a[1:])
	fmt.Println(a[:3])
	fmt.Println(len(a))
	fmt.Println(cap(a))
	var number []int
	if number==nil {
		fmt.Println("null")
	}
	a = append(a,0,1,23)
	//for i:=0;i<len(a) ; i++ {
	//	fmt.Println(i)
	//	fmt.Println(a[i])
	//
	//}
	numbers1 := make([]int, len(a), cap(a)*2)
	fmt.Println(numbers1)
}
