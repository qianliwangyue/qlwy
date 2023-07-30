package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"
)
func main(){
	http.HandleFunc("/",log(Index))
	http.ListenAndServe(":8087",nil)
}
func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"hello index!")
}
func log(h http.HandlerFunc) http.HandlerFunc{
	return  func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("  time: ",time.Now().String(),"|","handlerfunc:",runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name())
		h(w,r)
	}
}
/*
func log2(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w,r)
	})
}
*/