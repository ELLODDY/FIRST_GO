package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Словарь для преобразования римских цифр в арабские
var romanToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// Список для преобразования арабских чисел в римские
var arabicToRomanList = []struct {
	value int
	roman string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Функция для преобразования арабского числа в римское
func arabicToRomanFunc(num int) string {
	var roman string
	for _, entry := range arabicToRomanList {
		for num >= entry.value {
			roman += entry.roman
			num -= entry.value
		}
	}
	return roman
}

// Функция для преобразования римского числа в арабское
func romanToArabicFunc(roman string) (int, bool) {
	value, exists := romanToArabic[strings.ToUpper(roman)]
	return value, exists
}

// Функция для проверки, является ли строка арабским числом от 1 до 10
func isArabic(s string) (int, bool) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	if num < 1 || num > 10 {
		return 0, false
	}
	return num, true
}

// Функция для выполнения арифметической операции
func performOperation(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль невозможно")
		}
		return a / b, nil // Целочисленное деление
	default:
		return 0, fmt.Errorf("недопустимый оператор: %s", operator)
	}
}

// Функция для обработки ошибок и завершения программы
func panicError(message string) {
	fmt.Printf("Ошибка: %s\n", message)
	os.Exit(1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение (или 'exit' для выхода): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panicError("Не удалось прочитать ввод")
		}

		input = strings.TrimSpace(input)

		// Проверка на команду выхода
		if strings.EqualFold(input, "exit") || strings.EqualFold(input, "quit") {
			fmt.Println("Выход из программы. До свидания!")
			break
		}

		// Разбираем входную строку
		parts := strings.Fields(input)
		if len(parts) < 3 {
			panicError("Строка не является математической операцией.")
		} else if len(parts) > 3 {
			panicError("Формат математической операции не удовлетворяет заданию.")
		}

		operand1, operator, operand2 := parts[0], parts[1], parts[2]

		// Проверяем системы счисления
		aArabic, aIsArabic := isArabic(operand1)
		bArabic, bIsArabic := isArabic(operand2)

		aRoman, aIsRoman := romanToArabicFunc(operand1)
		bRoman, bIsRoman := romanToArabicFunc(operand2)

		var numberSystem string
		var a, b int

		if aIsArabic && bIsArabic {
			numberSystem = "arabic"
			a, b = aArabic, bArabic
		} else if aIsRoman && bIsRoman {
			numberSystem = "roman"
			a, b = aRoman, bRoman
		} else {
			panicError("Используются одновременно разные системы счисления или неверные числа.")
		}

		// Проверяем диапазон чисел (1-10)
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panicError("Числа должны быть от 1 до 10 включительно.")
		}

		// Выполняем операцию
		result, err := performOperation(a, b, operator)
		if err != nil {
			panicError(err.Error())
		}

		// Обработка результата
		if numberSystem == "roman" {
			if result < 1 {
				panicError("Результат в римской системе должен быть положительным.")
			}
			resultStr := arabicToRomanFunc(result)
			fmt.Printf("Результат: %s\n", resultStr)
		} else {
			fmt.Printf("Результат: %d\n", result)
		}

		fmt.Println() // Пустая строка для разделения вычислений
	}
}
