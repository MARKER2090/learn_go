package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	point := []string{"/hello", "/users/name2", "/views/*.html", "/order?id=234252"}[3]
	//creat the request
	request, _ := http.NewRequest("GET", "http://127.0.0.1:8082"+point, nil)

	response, _ := client.Do(request)
	if response.StatusCode != 400 {
		//connected
		body := response.Body
		defer response.Body.Close()
		wenben, _ := io.ReadAll(body)
		fmt.Println(string(wenben))
	}
}
