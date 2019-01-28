package main

import "fmt"

func main() {
	var i int = 5
	fmt.Printf("%d 的阶乘是 %d\n", i, actorial(uint64(i)))
}
func actorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * actorial(n-1)
		return result
	}
	return 1
}



