package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	go test()
	time.Sleep(time.Second)
}

func sayHello() {

	for i := 0; i < 10; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}
}

func test() {
	//myMap := make(map[int]string)
	/*
	*协程中的异常处理
	 */
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test 发错异常")
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}
