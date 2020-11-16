package main

import "fmt"

func main(){
	var param interface{}
	param = 66
	param_type := fmt.Sprintf("%T",param)
	if param_type == "int"{
		fmt.Println("整形")
	}
}