package main

import "fmt"

func divideAndRemainder(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("На ноль делить нельзя")
	}
	quotient := a / b  // результат деления
	remainder := a % b // остаток от деления
	return quotient, remainder, nil
}

func main() {
	var a, b int
	fmt.Print("Введите делимое и делитель: ")
	fmt.Scan(&a, &b)

	quotient, remainder, err := divideAndRemainder(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Частное: %d, Остаток: %d\n", quotient, remainder)
}
