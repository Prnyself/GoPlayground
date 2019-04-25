package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func main() {
	h := md5.New()
	str := `{"quiet": false}`
	h.Write([]byte(str)) // 需要加密的字符串为 {"quiet": false}
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	// 利用base64进行加密
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(cipherStr)) // 输出加密结果
}
