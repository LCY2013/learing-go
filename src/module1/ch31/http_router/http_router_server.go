package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "hello go")
}

func goTime(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	t := time.Now()
	timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t) + params.ByName("name")
	writer.Write([]byte(timeStr))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/time/:name", goTime)
	log.Fatal(http.ListenAndServe(":8080", router))
}
