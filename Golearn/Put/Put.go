package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://www.shirdon.com/comment/add"
	payload := strings.NewReader("{\"userId\":1,\"articleId\":1,\"comment\":\"这是一条评论\"}")
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

//"io/ioutil" has been deprecated since Go 1.19: As of Go 1.16, the same functionality is now provided by package io or package os, and those implementations should be preferred in new code. See the specific function documentation for details.  (SA1019)
