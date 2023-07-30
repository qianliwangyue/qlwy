package main

import (
	"fmt"
	"log"
	"net/http"
)

type WelcomeHandler struct {
	Language string
}

// 定义一个ServerHttp()方法，以实现Handler接口
func (h WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, h.Language)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/cn", WelcomeHandler{Language: "欢迎摆烂人光顾！"})
	mux.Handle("/en", WelcomeHandler{Language: "Wekcome to my world! Let's go Web!"})

	server := &http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
