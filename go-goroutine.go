package main

import (
	"fmt"
)

func main() {

	c:=make(chan bool)
	go func() {
		fmt.Println("echo")
		c<-true
	}()
	<-c

}

