package main

import (
	"errors"
	"fmt"
)

// 1. Создание структуры Employee
type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
}

// Метод Info() с возвратом данных о сотруднике
func (e Employee) Info() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d, Должность: %s, Зарплата: %.2f",
		e.Name, e.Age, e.Position, e.Salary)
}

// 2. Добавление сотрудника в срез
func AddEmployee(employees []Employee, newEmployee Employee) []Employee {
	return append(employees, newEmployee)
}

// Интерфейс
type Printable interface {
	Print()
}

// Реализация интерфейса Printable для Employee
func (e Employee) Print() {
	fmt.Println(e.Info())
}

// 3. Map: средняя зарплата по должностям
func UpdateAverageSalaries(employees []Employee) map[string]float64 {
	positionSalaries := make(map[string][]float64)
	for _, e := range employees {
		positionSalaries[e.Position] = append(positionSalaries[e.Position], e.Salary)
	}

	avgSalaries := make(map[string]float64)
	for pos, salaries := range positionSalaries {
		sum := 0.0
		for _, s := range salaries {
			sum += s
		}
		avgSalaries[pos] = sum / float64(len(salaries))
	}
	return avgSalaries
}

// 5. Функция поиска по имени
func FindEmployeeByName(employees []Employee, name string) (Employee, error) {
	for _, e := range employees {
		if e.Name == name {
			return e, nil
		}
	}
	return Employee{}, errors.New("сотрудник не найден")
}

// 5. Функция для среднего возраста
func AverageAge(employees []Employee) float64 {
	if len(employees) == 0 {
		return 0
	}
	sum := 0
	for _, e := range employees {
		sum += e.Age
	}
	return float64(sum) / float64(len(employees))
}

func main() {
	// Создаём сотрудников
	e1 := Employee{"Алиса", 25, "Заместитель директора", 5000}
	e2 := Employee{"Боб", 30, "Директор", 7000}
	e3 := Employee{"Роман", 28, "Заместитель директора", 5500}

	// Хранение в срезе
	var employees []Employee
	employees = AddEmployee(employees, e1)
	employees = AddEmployee(employees, e2)
	employees = AddEmployee(employees, e3)

	// Вывод всех сотрудников через интерфейс
	fmt.Println("Список сотрудников:")
	for _, e := range employees {
		var p Printable = e
		p.Print()
	}

	// Средний возраст
	fmt.Printf("\nСредний возраст: %.2f\n", AverageAge(employees))

	// Средняя зарплата по должностям
	fmt.Println("\nСредняя зарплата по должностям:")
	avgSalaries := UpdateAverageSalaries(employees)
	for pos, avg := range avgSalaries {
		fmt.Printf("%s -> %.2f\n", pos, avg)
	}

	// Поиск сотрудника
	fmt.Println("\nПоиск сотрудника 'Алиса':")
	found, err := FindEmployeeByName(employees, "Алиса")
	if err != nil {
		fmt.Println(err)
	} else {
		found.Print()
	}
}
