package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type respBody struct {
	Args struct {
		T string `json:"t"`
	} `json:"args"`
	Data  string `json:"data"`
	Files struct {
	} `json:"files"`
	Form struct {
	} `json:"form"`
	Headers struct {
		Accept                  string `json:"Accept"`
		AcceptEncoding          string `json:"Accept-Encoding"`
		AcceptLanguage          string `json:"Accept-Language"`
		Host                    string `json:"Host"`
		UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
		UserAgent               string `json:"User-Agent"`
	} `json:"headers"`
	JSON   interface{} `json:"json"`
	Method string      `json:"method"`
	Origin string      `json:"origin"`
	URL    string      `json:"url"`
}

func main() {
	reqURL := "http://httpbin.org/anything?t=123"
	//reqURL := "http://www.baidu.com"
	reqBody := &bytes.Buffer{}
	req, err := http.NewRequest("GET", reqURL, reqBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	HTTPClient := http.DefaultClient

	resp, err := HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	response, err := ioutil.ReadAll(resp.Body)

	res := &respBody{}
	fmt.Println(json.Unmarshal(response, res))
	fmt.Printf("res: %+v\n", res)
}
