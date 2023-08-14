package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "http://www.shirdon.com/comment/add"
	payload := strings.NewReader("{\"userId\":1,\"articleId\":1,\"comment\":\"这是一条评论\"}")
	req, _ := http.NewRequest("DELETE", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
