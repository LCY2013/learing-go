package easyjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"fufeng",
		"age":18
	},
	"job_info":{
		"skills":["java","go","c","c++"]
	}
}`

//测试go内置json解析
func TestEmbeddedJsonFormat(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*e)

	if ret, err := json.Marshal(e); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(ret))
	}
}

//测试easyJson解析
func TestEasyJsonFormat(t *testing.T) {
	e := new(Employee)
	err := e.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*e)
	}
	if ret, err := e.MarshalJSON(); err == nil {
		fmt.Println(string(ret))
	} else {
		fmt.Println(err)
	}
}

//BenchMark 比较两个方式的性能 go tool -bench=.
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	//b.StopTimer()
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
	//b.StopTimer()
}
