package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	fmt.Println("Начнём работу с данными!")

	var filename string

	// Делаем ввод
	fmt.Print("Укажите имя файла: ")
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		filename = scan.Text()
	} else {
		fmt.Println("Возникла ошибка обработки или чтения файла.", scan.Err())
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	var (
		set   Set
		stack Stack
		queue Queue
		table = HashTable{
			data:     make([]*NodeHT, 1000),
			capacity: 1000,
			mu:       sync.Mutex{},
		}
	)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Множество")
		fmt.Println("2. Стек")
		fmt.Println("3. Очередь")
		fmt.Println("4. Хэш-таблица")
		fmt.Println("5. Выход и завершение работы")

		if !scanner.Scan() {
			break
		}

		choice := scanner.Text()

		switch choice {
		case "1":
			menuSet(&set, scanner, filename)
		case "2":
			menuStack(&stack, scanner, filename)
		case "3":
			menuQueue(&queue, scanner, filename)
		case "4":
			menuHashTable(&table, scanner, filename)
		case "5":
			fmt.Println("Выход из программы.")
			os.Exit(0)
		default:
			fmt.Println("Указанной функции не существует")
		}
	}
}
