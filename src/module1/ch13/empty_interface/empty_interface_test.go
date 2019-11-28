package empty_interface

import (
	"fmt"
	"testing"
)

func DetectionType(i interface{}) {
	//if val,ok := i.(int);ok{
	//	fmt.Println("int type , value = ",val)
	//	return
	//}
	//if val,ok := i.(string);ok{
	//	fmt.Println("string type , value = ",val)
	//	return
	//}
	//fmt.Println("unknown type")
	switch val := i.(type) {
	case int:
		fmt.Println("int type , value = ", val)
	case string:
		fmt.Println("string type , value = ", val)
	default:
		fmt.Println("unknown type")
	}
}

func TestAssertVarType(t *testing.T) {
	DetectionType(1)
	DetectionType("1")
}

//单一方法接口原则，多方法用接口组合
type Reader interface {
	Reader(src string) []byte
}

type Writer interface {
	Writer(desc []byte) bool
}

type ReaderWriter interface {
	Reader
	Writer
}
