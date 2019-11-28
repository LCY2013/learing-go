package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

//获取变量的类型
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Printf("Integer,%T \n", v)
	case reflect.Float32, reflect.Float64:
		fmt.Printf("Float,%T \n", v)
	case reflect.String:
		fmt.Printf("string,%T \n", v)
	default:
		fmt.Printf("unknown,%T \n", v)
	}
}

func TestCheckType(t *testing.T) {
	var v = 10
	CheckType(&v)
}

func TestTypeAndValue(t *testing.T) {
	var v int64 = 10
	t.Log(reflect.TypeOf(v), reflect.ValueOf(v))
	t.Log(reflect.ValueOf(v).Type())
}

type Employee struct {
	Id   string
	Name string `format:"normal"` //给字段标记一个属性的Tag，struct Tag
	Age  int
}

func (e *Employee) UpdateAge(n int) {
	e.Age = n
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{Id: "12", Name: "fufeng", Age: 18}
	//通过反射获取字段名称为Name的字段值和字段类型
	t.Logf("Name : value(%[1]v),Name : type(%[1]T)",
		reflect.ValueOf(*e).FieldByName("Name"))
	//通过类型反射获取字段Name的是否存在
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); ok {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	} else {
		t.Log("get field 'Name' is unSuccess")
	}
	//通过反射调用结构体的方法
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(17)})
	t.Log("update age : ", e)
}

type Consumer struct {
	Id   string
	Name string
	Age  int
}

//深度比较
func TestDeepEqual(t *testing.T) {
	//定义两组map，在go语言中map只能与nil比较，所以我们需要deepEqual
	map1 := map[int]string{1: "1", 2: "2", 3: "3"}
	map2 := map[int]string{1: "1", 2: "2", 3: "3"}
	t.Log(reflect.DeepEqual(map1, map2))

	//定义一组数组
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{1, 2, 3, 4, 6}
	t.Log(reflect.DeepEqual(array1, array2))

	//定义一组自定义类型
	c1 := Consumer{Id: "1", Name: "c1", Age: 1}
	c2 := Consumer{Id: "1", Name: "c1", Age: 2}
	t.Log(reflect.DeepEqual(c1, c2))
}
