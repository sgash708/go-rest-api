package main

import (
	"io"
	"log"
	"net/http"
)

// main メイン処理
func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", Hello)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func Hello(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "Hello from handler!\n")
}
