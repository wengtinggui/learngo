package main

import "fmt"



func main() {
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Printf("key=%d\n",i)
	}
	fmt.Println(sum)
}
