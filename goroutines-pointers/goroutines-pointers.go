package main

import (
	"fmt"
	"time"
)

func Pointers(palavra *string) {
	for {
		fmt.Println(&palavra)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {
	var test string = "oi"
	var pointer *string = &test

	go Pointers(pointer)
	time.Sleep(time.Second)

	*pointer = "oi-2"
	time.Sleep(3 * time.Second)
}
