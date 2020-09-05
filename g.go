package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("你好 golang", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器减一
}

func main() {
	wg.Add() //协程计数器加一

	go test() //开启一个协程(就是用户级别的方法)
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait（）
	fmt.Println("主线程退出")

}
