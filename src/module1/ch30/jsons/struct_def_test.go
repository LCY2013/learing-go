package jsons

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

func TestJsonFormat(t *testing.T) {
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
