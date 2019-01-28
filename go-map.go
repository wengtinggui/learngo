package main

import "fmt"

func main() {
	var country map[string]string
	country = make(map[string]string)
	country["china"] = "中国"
	country["m"] = "美国"
	country["a"] = "澳大利亚"
	for c := range country {
		fmt.Println(c)
	}
	//删除
	delete(country,"m")
	delete(country,"a")
	for c := range country {
		fmt.Println(c)
	}
	var a map[int]int
	a=make(map[int]int)
	a[2]=1
	a[1]=2
	for c := range a {
		fmt.Println(c)
	}

}
