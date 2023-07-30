package main

import (
	"fmt"
)

func Sendmessage(user string) chan string {
	message := make(chan string, 500)
	go func() {
		message <- fmt.Sprintf("Hi %s,welcome to our site!", user)
	}()
	return message
}
func main() {
	name1 := Sendmessage("name1")
	name2 := Sendmessage("name2")
	fmt.Println(<-name1)
	fmt.Println(<-name2)
}
