package main

import (
	"fmt"
)

type Shower interface {
	Show() string
}

type ShowSct struct {
	Name string
}

func (s *ShowSct) Show() string {
	return s.Name
}

type AnotherShowSct struct {
	NameAlias string
}

func (s *AnotherShowSct) Show() string {
	return s.NameAlias
}

// type assertion for specific dynamic type, means "contain"
func AssertShowSct(v interface{}) string {
	if sct, ok := v.(*ShowSct); ok {
		return sct.Show()
	}
	return "v do not contains *ShowSct, dynamically"
}

// type assertion for interface
func AssertShowerInterface(v interface{}) string {
	if shower, ok := v.(Shower); ok {
		return shower.Show()
	}
	return "v do not implements shower interface"
}

func main() {
	sct := &ShowSct{Name: "Optic Valley Coder"}
	var shower Shower
	shower = sct
	fmt.Println(AssertShowSct(shower)) // Optic Valley Coder

	anoSct := &AnotherShowSct{NameAlias: "Another Optic Valley Coder"}
	shower = anoSct
	fmt.Println(AssertShowSct(shower)) // v do not contains *ShowSct, dynamically

	nsct := ShowSct{Name: "Not an Optic Valley Coder"}
	fmt.Println(AssertShowerInterface(nsct)) // v do not implements shower interface, cuz only *ShowSct has Show method

	fmt.Println(AssertShowerInterface(sct)) // Optic Vally Coder
}
