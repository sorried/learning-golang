package main

import "fmt" 

func main() {
	var name string // пустое, записать

	fmt.Printf("Как тебя зовут? ")
	fmt.Scan(name)

	var age int // сколько лет

	fmt.Printf("Сколько тебе лет? ")
	fmt.Scan(age)

	fmt.Printf("Привет %s!", name, "Тебе %d лет", age)
}
