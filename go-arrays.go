package main

import "fmt"

func main() {
	var n [10]int
	for i := 0; i < 10; i++ {
		n[i] = i + 1
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("数组N[%d]=%d\n", j, n[j])
	}
	var a=[]int{1,2,3,4,5}
	setArray(a)

}
func setArray(a []int)  {
	fmt.Println("len",len(a))

}
