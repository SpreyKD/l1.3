package main

import (
	"fmt"
	"time"
)

func main() {
	WorkersNum := 5 //кол-во воркеров
	chanel := make(chan int)

	for i := 1; i <= WorkersNum; i++ {
		go workers(i, chanel)
	}

	perebor := 1
	for {
		chanel <- perebor
		perebor++
		time.Sleep(200 * time.Millisecond)
	}
}

func workers(id int, ch <-chan int) {
	for val := range ch {
		fmt.Printf("номер %d значение %d\n", id, val)
	}
}
