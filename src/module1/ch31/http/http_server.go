package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write([]byte("hello go"))
		fmt.Fprintf(writer, "hello go")
	})
	http.HandleFunc("/time/", func(writer http.ResponseWriter, request *http.Request) {
		timeNow := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", timeNow)
		writer.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080", nil)
}
