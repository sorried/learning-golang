package main

import "fmt"

func main() {
	const name = "Max"
	var love string = "cats"
	var age int = 20
	sex := "male"

	fmt.Println("2025")
	fmt.Println("name", name)
	fmt.Println("love: ", love)
	fmt.Println("age: ", age)
	fmt.Println("sex: ", sex)

	fmt.Println("2026")
	love = "cats"
	age = 21
	sex = "male"

	fmt.Println("name", name)
	fmt.Println("love: ", love)
	fmt.Println("age: ", age)
	fmt.Println("sex: ", sex)
}
