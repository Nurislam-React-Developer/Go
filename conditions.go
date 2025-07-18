package main

import "fmt"

func main() {
	age := 20

	if age >= 18 {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}

	// if с else if
	score := 75

	if score >= 90 {
		fmt.Println("Оценка: Отлично")
	} else if score >= 70 {
		fmt.Println("Оценка: Хорошо")
	} else {
		fmt.Println("Оценка: Удовлетворительно")
	}
}
