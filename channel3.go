package main

import "fmt"

func main() {
	syncchan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		syncchan <- i
	}

	close(syncchan)
	var res int
	for i := 0; i < 10; i++ {
		res = <-syncchan
		fmt.Println(res)
		fmt.Println(len(syncchan))
		//syncchan <- res
		//if res == 5 {
		//	close(syncchan)
		//}
	}
}
