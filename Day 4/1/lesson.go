package main

import "fmt"

func main() {
	// Способ 1: Сразу заполнить значениями
	// Ключ — строка (название софта), значение — целое число (порт)
	ports := map[string]int{
		"http":  80,
		"https": 443,
		"ssh":   22,
	}

	// Способ 2: Создать пустую карту через make
	// и положить данные потом
	users := make(map[string]string)
	users["admin"] = "root"
	users["cldoff"] = "arch"

	fmt.Println("Порт SSH:", ports["ssh"])
	fmt.Println("Юзер sorried:", users["sorried"])
	scores := make(map[string]int)

	// Добавить или обновить
	scores["max"] = 100
	scores["max"] = 150 // Перезаписали старое значение

	// Удалить ключ с помощью встроенной функции delete
	delete(scores, "max")
	port, exists := ports["dns"]

	if exists {
		fmt.Println("Порт найден:", port)
	} else {
		fmt.Println("Такого сервиса нет в базе!")
	}
	for service, port := range ports {
		fmt.Printf("Сервис: %s -> Порт: %d\n", service, port)
	}
	}
