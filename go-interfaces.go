package main

import "fmt"

type Phone interface {
	call()
}
type nokiaPhone struct {

}

func (nokiaPhone nokiaPhone) call()  {
	fmt.Println("nokiaPhone call")
}

type iphone struct {

}

func (iphone iphone) call() {
	fmt.Println("iphone call")

}
func main() {
	var phone Phone
	phone = new (nokiaPhone)
	phone.call()
	phone = new (iphone)
	phone.call()
}



