package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Список проектов (ключ — имя для отображения, значение — путь к файлу или команда)
var projects = map[int]struct {
	name    string
	command string
}{
	1: {"Перцептрон (четность чисел)", "go run ./1.0_oddeven/main.go"},
	// Добавляй сюда новые проекты по мере их появления
}

func main() {
	fmt.Println("Добро пожаловать в ML & Go!")
	fmt.Println("Доступные проекты:")

	// Показываем меню
	for id, proj := range projects {
		fmt.Printf("%d. %s\n", id, proj.name)
	}
	fmt.Println("0. Выход")
	fmt.Print("Выберите проект (введите номер): ")

	// Читаем выбор пользователя
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	// Преобразуем ввод в число
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 0 || choice > len(projects) {
		fmt.Println("Неверный выбор. Завершение программы.")
		return
	}

	// Выход при выборе 0
	if choice == 0 {
		fmt.Println("До встречи!")
		return
	}

	// Получаем команду для выбранного проекта
	proj, exists := projects[choice]
	if !exists {
		fmt.Println("Проект не найден.")
		return
	}

	// Разбиваем команду на части (например, "go run perceptron.go" → ["go", "run", "perceptron.go"])
	parts := strings.Fields(proj.command)
	if len(parts) < 2 {
		fmt.Println("Ошибка в команде проекта.")
		return
	}

	// Запускаем проект
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = "."          // Устанавливаем рабочую директорию (текущая папка ML & Go)
	cmd.Stdout = os.Stdout // Перенаправляем вывод в консоль
	cmd.Stderr = os.Stderr // Перенаправляем ошибки тоже

	fmt.Printf("\nЗапускаю проект: %s\n", proj.name)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Ошибка запуска: %v\n", err)
		return
	}

	fmt.Println("\nПроект завершен.")
}
