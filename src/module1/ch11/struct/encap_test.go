package encap__test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

//go语言创建结构体对象
func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"1", "fufeng", 20}
	e1 := Employee{Name: "fufeng1", Age: 18}
	e2 := new(Employee)
	e2.Id = "3"
	e2.Name = "fufeng3"
	e2.Age = 17
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e type is %T", e)
	t.Logf("e2 type is %T", e2)
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id %s - Name %s - Age %d", e.Id, e.Name, e.Age)
}

func (e *Employee) Strings() string {
	fmt.Printf("* Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id %s / Name %s / Age %d", e.Id, e.Name, e.Age)
}

func TestEmployeeString(t *testing.T) {
	e := Employee{"12", "fufeng12", 17}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	t.Log(e.Strings())
}
