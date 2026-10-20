// Срезы (Slices) — списки данных в Go

package main

import (
	"fmt"
)

func main() {
	// квадрат. скопки для элементов (чтоб был сразу пустой)
	foods := []string{"apple", "cake", "burger", "pizza"}

	var what []string

	what = append(what, "ничего")
	what = append(what, "и ничего")

	fmt.Println("еда ", foods)
	fmt.Println("принесли", what)

	fmt.Println(foods[0])
	fmt.Println(foods[1])

	for i, food := range foods {
		fmt.Printf("индекс %d: название %s\n", i, food)
	}
}
