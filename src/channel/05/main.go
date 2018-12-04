package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello " + fmt.Sprintf("%d", i)
	}

outFor:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从intChan读取数据%s\n", v)
		default:
			fmt.Println("ok")
			break outFor
		}
	}
}
