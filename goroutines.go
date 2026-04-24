package main

import "time"

func printMessage(text string) {
	for i := 0; i < 5; i++ {
		println(text)
		time.Sleep(time.Second) // sleep for 1 second
	}

}
