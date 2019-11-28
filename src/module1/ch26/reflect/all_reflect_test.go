package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

//实现给任意对象填充字段的值
func fillBySettings(sts interface{}, settings map[string]interface{}) error {
	//func (v Value) Elem() Value
	//Elem returns the value that the interface v contains or that the pointer
	//It panics if v's Kind is not interface or Ptr
	//It return the zero Value if v is nil
	if reflect.TypeOf(sts).Kind() != reflect.Ptr {
		//Elem() 获取指针指向的值
		if reflect.TypeOf(sts).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to struct type")
		}
	}

	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		//结构体中是否存在该字段
		if field, ok = reflect.ValueOf(sts).Elem().Type().FieldByName(k); !ok {
			continue
		}
		//如果结构体字段相同，就开始设值
		if field.Type == reflect.TypeOf(v) {
			//设值
			valueRet := reflect.ValueOf(sts).Elem()
			valueRet = valueRet.FieldByName(k)
			valueRet.Set(reflect.ValueOf(v))
		}
	}
	return nil
}

//测试所以字段填充
func TestFillBySettings(t *testing.T) {
	//定义一个填充结构体的map对象
	settings := map[string]interface{}{"Id": "11", "Name": "fufeng", "Age": 12}

	e := Employee{}
	fmt.Println("before:", e)
	if error := fillBySettings(&e, settings); error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("after:", e)
	}

	c := new(Consumer)
	fmt.Println("before:", c)
	if error := fillBySettings(c, settings); error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("after:", c)
	}
}
