package tool

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

const jsonStr = `{
	"basic_info":{
		"name":"fufeng",
		"age":18
	},
	"job_info":{
		"skills":["java","go","c","c++"]
	}
}`

//var jsonStr = `
//	{"basic_info":{"name":"fufeng","age":18},"job_info":{"skills":["java","go","c","c++"]}}
//`

func ProcessOld(e *Employee) {
	var ret string
	for i := 0; i < 100; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			fmt.Println(err)
		}
		if v, err := json.Marshal(e); err != nil {
			fmt.Println(err)
		} else {
			ret += string(v)
		}
	}
}

//BenchMark 比较两个方式的性能 go tool -bench=.
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		ProcessOld(e)
	}
	b.StopTimer()
}

func Process(e *Employee) string {
	var buf strings.Builder
	for i := 0; i < 100; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			fmt.Println(err)
		}
		if v, err := e.MarshalJSON(); err != nil {
			fmt.Println(err)
		} else {
			buf.WriteString(string(v))
		}
	}
	return buf.String()
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		fmt.Println(Process(e))
	}
	b.StopTimer()
}
