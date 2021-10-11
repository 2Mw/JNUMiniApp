package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// 请求体构建
	link, _ := url.Parse("http://httpbin.org/get?Name=1")
	params := url.Values{}
	params.Set("Age", "18")
	link.RawQuery += "&" + params.Encode() //	添加GET请求体

	// 发送请求

	client := http.Client{}
	req, err := http.NewRequest("GET", link.String(), nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "Chrome")

	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println("网络出现错误")
		fmt.Println(err)
	}
	fmt.Println("请求成功：")
	fmt.Println("请求协议", rsp.Proto)
	fmt.Println("响应码", rsp.StatusCode)
	fmt.Println("响应头：")

	for i, v := range rsp.Header {
		fmt.Println("\t", i, v)
	}

	defer rsp.Body.Close() //	等待请求关闭

	// 处理响应体
	body, err := io.ReadAll(rsp.Body)
	// JSON 数据处理

	type MSG struct {
		Args struct {
			Age  string `json:"Age"`
			Name string `json:"Name"`
		} `json:"args"`
		Headers struct {
			AcceptEncoding string `json:"Accept-Encoding"`
			Host           string `json:"Host"`
			UserAgent      string `json:"User-Agent"`
			XAmznTraceId   string `json:"X-Amzn-Trace-Id"`
		} `json:"headers"`
		Origin string `json:"origin"`
		Url    string `json:"url"`
	}

	dec := json.NewDecoder(strings.NewReader(string(body)))
	var m MSG
	if err = dec.Decode(&m); err == io.EOF {
		fmt.Println("JSON 内容为空")
	} else if err != nil {
		log.Fatal("Decode JSON Data Error!")
	}
	fmt.Println("UA:", m.Headers.UserAgent)

}
