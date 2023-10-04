package main

import (
	"bufio"
	"fmt"
)

func menuSet(set *Set, scanner *bufio.Scanner, filename string) {
	for {
		fmt.Println("\nВыберите действие для множества:")
		fmt.Println("1. Добавление (SADD)")
		fmt.Println("2. Удаление (SREM)")
		fmt.Println("3. Чтение (SISMEMBER)")
		fmt.Println("4. Отмена")

		if !scanner.Scan() {
			break
		}

		action := scanner.Text()

		switch action {
		case "1":
			fmt.Print("Добавим элемент: ")
			if scanner.Scan() {
				element := scanner.Text()
				set.SADD(filename, element)
			}
		case "2":
			fmt.Print("Удалим элемент: ")
			if scanner.Scan() {
				element := scanner.Text()
				set.SREM(filename, element)
			}
		case "3":
			fmt.Print("Читаем элемент: ")
			if scanner.Scan() {
				element := scanner.Text()
				if set.SISMEMBER(filename, element) {
					fmt.Println(element, "присутствует в множестве")
				} else {
					fmt.Println(element, "не найден в множестве")
				}
			}
		case "4":
			return
		default:
			fmt.Println("Выбор некорректен, попробуйте снова, пожалуйста.")
		}
	}
}

func menuStack(stack *Stack, scanner *bufio.Scanner, filename string) {
	for {
		fmt.Println("\nВыберите действие для стека:")
		fmt.Println("1. Добавление (SPUSH)")
		fmt.Println("2. Удаление (SPOP)")
		fmt.Println("3. Отмена.")

		if !scanner.Scan() {
			break
		}

		action := scanner.Text()

		switch action {
		case "1":
			fmt.Print("Добавим элемент в Стек: ")
			if scanner.Scan() {
				element := scanner.Text()
				stack.SPUSH(filename, element)
			}
		case "2":
			_, err := stack.SPOP(filename)
			if err != nil {
				fmt.Println("В Стеке нет элементов.")
			}
		case "3":
			return
		default:
			fmt.Println("Выбор некорректен, попробуйте снова, пожалуйста.")
		}
	}
}

func menuQueue(queue *Queue, scanner *bufio.Scanner, filename string) {
	for {
		fmt.Println("\nВыберите действие для очереди:")
		fmt.Println("1. Добавление (QPUSH)")
		fmt.Println("2. Удаление (QPOP)")
		fmt.Println("3. Отмена.")

		if !scanner.Scan() {
			break
		}

		action := scanner.Text()

		switch action {
		case "1":
			fmt.Print("Добавим элемент в очередь: ")
			if scanner.Scan() {
				element := scanner.Text()
				queue.QPUSH(filename, element)
			}
		case "2":
			_, err := queue.QPOP(filename)
			if err != nil {
				fmt.Println("Очередь пуста.")
			}
		case "3":
			return
		default:
			fmt.Println("Выбор некорректен, попробуйте снова, пожалуйста.")
		}
	}
}

func menuHashTable(table *HashTable, scanner *bufio.Scanner, filename string) {
	for {
		fmt.Println("\nВыберите действие для хэш-таблицы:")
		fmt.Println("1. Добавление (HSET)")
		fmt.Println("2. Удаление (HDEL)")
		fmt.Println("3. Чтение (HGET)")
		fmt.Println("4. Отмена")

		if !scanner.Scan() {
			break
		}

		action := scanner.Text()

		switch action {
		case "1":
			fmt.Print("Добавим следующий ключ: ")
			if scanner.Scan() {
				key := scanner.Text()
				fmt.Print("Введите значение: ")
				if scanner.Scan() {
					value := scanner.Text()
					table.HSET(filename, key, value)
				}
			}
		case "2":
			fmt.Print("Ключ для удаления: ")
			if scanner.Scan() {
				key := scanner.Text()
				table.HDEL(filename, key)
			}
		case "3":
			fmt.Print("Ключ для чтения: ")
			if scanner.Scan() {
				key := scanner.Text()
				value, err := table.HGET(filename, key)
				if err != nil {
					fmt.Println("Ошибка:", err)
				} else {
					fmt.Println("Значение:", value)
				}
			}
		case "4":
			return
		default:
			fmt.Println("Выбор некорректен, попробуйте снова, пожалуйста.")
		}
	}
}
