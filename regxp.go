package main

import (
	"fmt"
	"regexp"
)

func main() {
	st := regexp.QuoteMeta(`\z`)
	fmt.Println(st)
	reg := "/str\\z"
	myReg, _ := regexp.Compile(reg)
	s := "/str"
	fmt.Printf("string:%s, reg:%s, %v\n", s, myReg.String(), myReg.MatchString(s))

	reg2 := `^/(?P<v0>[^/]+)/(?P<v1>[^/]+)$`
	myReg2, _ := regexp.Compile(reg2)
	s2 := "/str/act/"
	fmt.Println(myReg2.MatchString(s2))

	bucketRegStr := `^[a-z\d][a-z-\d]{4,61}[a-z\d]$`
	matchStr := "qingstor-bucket"
	matched, err := regexp.MatchString(bucketRegStr, matchStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("string: %s, match: %v", matchStr, matched)
}
