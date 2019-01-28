package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func main() {

	var c1 Circle
	c1.radius=100
	fmt.Println("圆的面积为:",c1.getArea())
	//函数作为值
	getSqr := func(x float64) float64{
		return math.Sqrt(x)
	}
	fmt.Println(getSqr(9))

}

func (c Circle) getArea() float64 {

	return 3.14 * c.radius * c.radius
}
