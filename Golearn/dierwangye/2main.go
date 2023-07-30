package main

import (
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,少年！"))
}

func main() {

	http.HandleFunc("/hello", sayhello)
	http.ListenAndServe(":8080", nil)

}
