package main

import (
	"fmt"
	"net/http"
)

func init() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world")
	})
}
