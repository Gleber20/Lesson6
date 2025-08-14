package main

import "fmt"

type Person struct {
	Name   string
	Salary float64
}

func main() {
	p := Person{
		Name:   "Алиса",
		Salary: 1999.9,
	}

	fmt.Println("Имя:", p.Name)
	fmt.Println("Зарплата:", p.Salary)
}
