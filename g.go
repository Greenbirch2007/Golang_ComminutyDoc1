package main

import (
	"fmt"
	"sync"
	"time"
)


var mutex sync.RWMutex

var wg sync.WaitGroup

//写的方法

func write(){
	mutex.Lock()//写的时候一个人来写
	fmt.Println("执行写操作")
	time.Sleep(time.Second*2)
	mutex.Unlock()
	wg.Done()
}


//读的方法

func read(){
	mutex.RLock() //读的时候多个人来读
	fmt.Println("----执行读操作")
	time.Sleep(time.Second*2)
	mutex.RUnlock()
	wg.Done()
}


func main(){


	//开启10个协程执行读操作
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go write()
	}
	//开启10个协程执行写操作
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go read()
	}





	wg.Wait()

}