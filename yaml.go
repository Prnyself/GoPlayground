package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YamlTest struct {
	A       map[string]string `yaml:"a2"`
	Account map[string]string `yaml:"acc"`
}

func main() {
	cYAML, err := ioutil.ReadFile("yaml_test.yaml")
	if err != nil {
		fmt.Println(err)
	}
	s := &YamlTest{}
	err = yaml.Unmarshal(cYAML, s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("s.Account nil?", s.Account == nil)
	fmt.Println(s.A["create"])
	fmt.Println(s.A["move"])
	fmt.Println(map[string]string{"a": "1", "b": "2"})
}
