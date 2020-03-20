package testmyself

import (
	"encoding/json"
	"fmt"
)

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/3/6 21:29
 */
type Action struct {
	Do  string `json:"do"`
	Say string `json:"say"`
}

type OutSide struct {
	Height int `json:"height"`
}

type people struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Action         // 匿名组合 相当于同一级
	Out    OutSide `json:"out"`
}

func StructTest() {
	var s = people{
		Name: "lihua",
		Age:  20,
		Action: Action{
			Do:  "walk",
			Say: "hello",
		},
		Out: OutSide{Height: 178},
	}
	out, _ := json.Marshal(s)
	fmt.Println(string(out))
}
