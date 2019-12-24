package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//encode()
	decode()
}

type Customer struct {
	Name string `json:"name"`
	IsMember bool `json:",string"`
	Goods []string `json:"goods"`
	Height float64 `json:"-"`
	email string
}

func encode() {
	customer := Customer{"jack", true, []string{"MacBook Pro", "Iphone", "Switch"}, 1.65, "a@qq.com"}
	fmt.Println(customer)

	result, err := json.MarshalIndent(customer, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(result))

	customerMap := make(map[string]interface{})
	customerMap["name"] = "pain"
	customerMap["isMember"] = true
	customerMap["goods"] = []string{"Cup", "Chair", "Bag"}
	customerMap["height"] = 1.68
	customerMap["email"] = "pain@163.com"

	fmt.Println(customerMap)

	result, err = json.MarshalIndent(customerMap, "", "    ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(result))
}

func decode() {
	customerStr := `{
  		"name": "jackSon",
  		"isMember": "true",
  		"goods": [
    		"MacBook Pro",
    		"Iphone",
    		"Switch",
			"Bag"
  		]
	}`
	var customer Customer

	err := json.Unmarshal([]byte(customerStr), &customer)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("customer: %+v\n", customer)

	var customerMap map[string]interface{}
	err = json.Unmarshal([]byte(customerStr), &customerMap)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("customerMap: %+v\n", customerMap)

	// 类型断言
	//var name string = customerMap["name"]
	//fmt.Println(name)
}