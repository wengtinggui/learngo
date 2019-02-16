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
	a[5]=1
	a[6]=2
	for i,c := range a {
		fmt.Println(i)
		fmt.Println(c)
	}
	b:=make(map[string]string)
	b["name"]="zangsan";
	b["age"]="11";
	fmt.Println(b)
	for i,c:=range b   {
		fmt.Println(i)
		fmt.Println(c)
	}

}
