package main

import (
	"fmt"
	"github.com/tabalt/gracehttp"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	err := gracehttp.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
