package main

import "fmt"

func main(){
	var param interface{}
	param = "66"
	switch param.(type) {
	case int:
		fmt.Println("整型")
	case  string:
		fmt.Println("字符串")
	default:
		fmt.Println("qit")
	}
}
