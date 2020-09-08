package main

import (
	"fmt"
	"sync"
	"time"
)

/*
goroutine和channel结合解决两个问题
1. 如果变化开启的协程数，4 6 9 如何简化算法
2. 想一遍统计一遍打印 统计和打印素数并行执行
单独的goroutine无法实现数据的共享

*/

var wg sync.WaitGroup
func putNum(intChan chan int){
	for i:= 2;i<12000;i++{
		intChan <- i
	}
	close(intChan)
	wg.Done()
}


func printPrime(primeChan chan int){
	for v := range primeChan{
		fmt.Println(v)
	}
	wg.Done()
}

func primeNum(intChan chan int,primeChan chan int,exitChan chan bool){
	for num := range intChan{
		var flag = true
		for i :=2;i<num;i++{
			if num %i ==0{
				flag=false
				break
			}
		}
		if flag {
			primeChan <- num //num是素数
		}

	}
	//要关闭primeChan
	//close(primeChan) //如果一个cahnnel关闭了就无法给这个channel发送数据了
	//什么时候关闭primeChan
	exitChan <- true
	wg.Done()
	//执行完一次后，给exitChan里面放入一条数据

}



func main(){
	start := time.Now().Unix()
	intChan := make(chan int,1000)
	primeChan :=make(chan int,1000)
	exitChan := make(chan bool,16) //标识primeChan close
	//存放数字的协程
	wg.Add(1)
	go putNum(intChan)
	for i:=0;i<16 ;i++  {
		wg.Add(1)
		//统计素数的协程
		go primeNum(intChan,primeChan,exitChan)
	}
	//打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)
	//判断exitChan是否存满
	wg.Add(1)
	go func() {
		for i:=0;i<16;i++{
			//取值
			<- exitChan
		}
		//关闭primeChan
		close(primeChan)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("执行完毕")
	end := time.Now().Unix()
	fmt.Println(end-start)

}