// карточка системы
package main

import "fmt"

func printSystemInfo() {
	fmt.Println("OS: Arch Linux")
	fmt.Println("Shell: zsh")
}

func main() {
	fmt.Println("проверка системы")

	printSystemInfo()

	fmt.Println("проверка окончена ")
}
