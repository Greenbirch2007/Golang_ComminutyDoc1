package main

import "fmt"

func main(){
	var param interface{}
	param=66
	_,yes := param.(float64)
	if yes{
		fmt.Println("是float64")
	}else{
		fmt.Println("no~")
	}
}
