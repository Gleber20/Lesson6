package main

import "fmt"

func operation(a, b int, op func(int, int) (int, int, error)) (int, int, error) {
	return op(a, b)
}

func divideAndRemainder(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("На ноль делить нельзя")
	}
	return a / b, a % b, nil
}

func main() {
	var a, b int
	fmt.Print("Введите делимое и делитель: ")
	fmt.Scan(&a, &b)

	quotient, remainder, err := operation(a, b, divideAndRemainder)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Частное: %d, Остаток: %d\n", quotient, remainder)
}
