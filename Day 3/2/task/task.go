package main

import (
	"fmt"
	"time"
)

func main() {
	wl := []string{"озу", "второй монитор", "кабеляяяяяяяя"}

	for i, wll := range wl {
		fmt.Printf("вишлист №%d: %s\n", i, wll)
	}

	time.Sleep(300 * time.Millisecond)

	wl = append(wl, "микросхемы")
	fmt.Printf("обновленный список %s\n", wl)
}
