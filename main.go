package main

import (
	"fmt"
	"time"
)

func add(palavra string) {
	for i := 0; i < 5; i++ {
		fmt.Println(palavra)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	add("Hello World")

	go add("Cheguei Brasil")
	time.Sleep(5 * time.Millisecond)

	fmt.Println("DONE")
}
