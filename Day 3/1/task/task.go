package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 5; i <= 5; i-- {
		fmt.Printf("круг номер %d\n", i)
		

	if i <= 1 {
		fmt.Println("полный газ!")
		break
	}
	time.Sleep(300 * time.Millisecond)
	/*
	time.Second — 1 секунда

	time.Millisecond — 1 миллисекунда

	time.Minute — 1 минута
	*/
	}
}
