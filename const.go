package main

import "fmt"

type SameSite int

const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
)

func main() {
	fmt.Printf("SameSiteDefaultMode=%d, SameSiteLaxMode=%d, SameSiteStrictMode=%d",
		SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode)
}
