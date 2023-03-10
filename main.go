package main

import (
	"io"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}


func main() {
	port := os.Getenv("PORT");
	fake := os.Getenv("PORT");
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":" + port, nil)
}