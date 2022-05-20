package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := http.Client{}

	// 创建一个 HTTP 请求
	req, err := http.NewRequest("POST", "http://163.com/", strings.NewReader("key = value"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "myClient")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// 读取服务器返回的内容
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	defer resp.Body.Close()
}
