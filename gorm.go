package main

import (
	"awesomeProject/model"
	"fmt"
)



func main() {
	app := model.CampusApp{}
	app = model.GetAppById(1)
	fmt.Println(app)
}
