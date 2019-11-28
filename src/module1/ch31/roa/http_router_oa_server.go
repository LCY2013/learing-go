package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Employee struct {
	Name string
	Age  int
}

var mem map[string]*Employee

func init() {
	mem = make(map[string]*Employee)
	mem["luo"] = &Employee{Name: "luo", Age: 18}
	mem["chun"] = &Employee{Name: "chun", Age: 17}
	mem["yun"] = &Employee{Name: "yun", Age: 16}
}

func index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "welcome to magic OA")
}

func find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	var (
		ret      *Employee
		jsonByte []byte
		ok       bool
		err      error
	)
	if ret, ok = mem[name]; !ok {
		writer.Write([]byte("{\"error\":\"" + name + "not found\"}"))
		return
	}
	if jsonByte, err = json.Marshal(ret); err != nil {
		writer.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err)))
		return
	}
	writer.Write([]byte(jsonByte))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/employee/:name", find)
	http.ListenAndServe(":8080", router)
}
