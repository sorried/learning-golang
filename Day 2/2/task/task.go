package main 

import ( // вызов > 1 библиотеки
	"fmt"
	"strings"
)

func main() {
	var distro string
	fmt.Printf("Какой у тебя дистрибутив? ")
	fmt.Scan(&distro)

	distro = strings.ToLower(distro) // в нижний регистр

	if distro == "arch" { 
		fmt.Println("i use arch btw")
	} else {
		fmt.Println("заебок")
	}
}

/*

чтобы сделать это в один if надо писать || (сравнить т.е.) 

*/
