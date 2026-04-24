package main

import (
	"fmt"
	"time"
)

func channel() {
	ch := make(chan int)

	go func() {
		ch <- 100 // send data
	}()

	value := <-ch // receive data

	fmt.Println(value)
}

func worker(ch chan string) {

	time.Sleep(2 * time.Second)

	ch <- "Work completed"
}

func worker2() {

	ch := make(chan string)

	go worker(ch)

	msg := <-ch

	fmt.Println(msg)
}
