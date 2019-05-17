package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	UrlStr := "/pool?a=1&b=%2&c=中国"
	l := url.PathEscape(UrlStr)
	fmt.Println(UrlStr, "pathEscape:", l)
	lu, err := url.PathUnescape(l)
	fmt.Println(l, "pathUnescape:", lu, err)

	lu2, err := url.PathUnescape("/pool?a=1&b=%252&c=%E4%B8%AD%E5%9B%BD")
	// 解析之后是  /pool?a=1&b=%2&c=中国
	fmt.Println("/pool?a=1&b=%252&c=%E4%B8%AD%E5%9B%BD pathUnescape:", lu2, err)

	fmt.Println(len(strings.Split(UrlStr, "?")))
	//res, err := url.ParseQuery(UrlStr)
	//fmt.Println(res)
	oriStr := "http://baidu.com/path/to/rules?metric&metric=1"
	queryMap, err := url.ParseQuery(strings.Split(oriStr, "?")[1])
	if err != nil {
		println(err)
	}
	fmt.Println("query:", queryMap)
	if v, ok := queryMap["metric"]; ok {
		fmt.Println("v:", v, "v1:", v[0], len(v))
		fmt.Printf("type of v: %T", v)
	}
	if v, ok := queryMap["nonexistence"]; ok {
		fmt.Println("v:", v, "v1:", v[0], len(v))
		fmt.Printf("type of v: %T", v)
	}
	//fmt.Println(len(queryMap.Get("metric")) == 0)
	//u, _ := url.Parse(oriStr)
}
