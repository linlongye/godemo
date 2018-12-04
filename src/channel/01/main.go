package main

import (
	"fmt"
)

func main() {
	var intChan = make(chan int, 3)

	fmt.Printf("intChan = %v\n", intChan)

	intChan <- 10
	num := 122
	intChan <- num

	fmt.Printf("channel len = %v cap = %v\n", len(intChan), cap(intChan))

	//遍历channel前需要关闭channel
	close(intChan)
	// 从管道读取数据
	for v := range intChan {
		fmt.Println("v = ", v)
	}
}
