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

func testStruct(){
	monster := Monster{
		Name: "Vincent",
		Age: 18000,
		Birthday: "2020-2-2",
		Sal: 1000000,
		Skill: "Skill",
	}

	content, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", content)
}

func testMap(){
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "Vincent"
	a["age"] = 199
	a["address"] = "Chengdu,Sichuan"

	content, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", content)

}

func testSlice(){
	var slice []map[string]interface{}
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "Vincent"
	a["age"] = 199
	a["address"] = "Chengdu,Sichuan"

	var b map[string]interface{}
	b = make(map[string]interface{})
	b["name"] = "Vincent"
	b["age"] = 199
	b["address"] = "Chengdu,Sichuan"

	slice = append(slice, a)
	slice = append(slice, b)


	content, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", content)
}

func main(){
	// 结构体, map, 切片的序列化
	testStruct()

	testMap()

	testSlice()
}
