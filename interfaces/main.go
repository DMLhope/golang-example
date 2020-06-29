//接口

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64  //面积
	perim() float64 //周长
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64 //半径
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

//接口这边仍然有很多疑惑只能等用的时候再加深理解。。。
//更详细的参考地址:https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
