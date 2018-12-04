package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)

	// 开启一个协程向intChan中放入数据
	go putNum(intChan)

	//开启四个协程，从intChan中取出数据，并判断是否为素数，如果是放入primeChan中
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	//遍历结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数 = %d\n", res)
	}
	fmt.Println("主线程退出")
	end := time.Now().UnixNano()
	fmt.Printf("总共耗时:%v", end-start)

}

func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//var num int
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		//判断是否是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该数不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("一个协程退出")
	exitChan <- true
}
