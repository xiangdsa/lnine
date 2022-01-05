package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http get by proxy
func GetByProxy(url string) (string, error) {
	request, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	resp, _ := client.Do(request)
	fmt.Println(resp)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body),nil
}

func main() {
	//proxy := "http://58.252.56.149:9000/"
	url := "https://www.daque.cn/softdown/43468.html"
	GetByProxy(url)
}