package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"xiaokeai"`
	Age int `json:"nianling"`
	Birthday string `json:"shengri"`
	Sal float64 `json:"gongzi"`
	Skill string `json:"jineng"`
}

func unmarshalStruct(){
	str := "{\"xiaokeai\":\"Vincent\",\"nianling\":18000,\"shengri\":\"2020-2-2\",\"gongzi\":1000000,\"jineng\":\"Skill\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(monster)
}

func main(){
	unmarshalStruct()
}
