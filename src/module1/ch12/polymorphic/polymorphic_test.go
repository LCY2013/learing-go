package polymorphic

import (
	"fmt"
	"testing"
)

type Programmer interface {
	SayHello()
}

type goProgrammer struct {
}

func (g *goProgrammer) SayHello() {
	fmt.Println("fmt.Println(\"hello fufeng\")")
}

type javaProgrammer struct {
}

func (java *javaProgrammer) SayHello() {
	fmt.Println("System.out.println(\"hello fufeng\")")
}

func sayHello(p Programmer) {
	p.SayHello()
}

//测试多态
func TestPolymorphic(t *testing.T) {
	g := new(goProgrammer)
	java := &javaProgrammer{}
	sayHello(g)
	sayHello(java)
}
