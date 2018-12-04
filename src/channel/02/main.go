package main

import "fmt"

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	//只读channel
	// var chanWriteOnly chan<-int

	//只写channe
	// var chanReadOnly <- chan int

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读取到数据 = %v\n", v)
	}

	//读取完成后
	exitChan <- true
	close(exitChan)
}
