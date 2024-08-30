package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func main() {
	fmt.Println("------------------------------Условия---------------------------------------- \n" +
		"Данный калькулятор умеет работать с целыми числами от 1 до 10 (в том числе арабскими), однако не одновременно.\n" +
		"Результат деления - целое число, а его остаток отбрасывается.\n" +
		"Арабские числа могут привести к отрицательному числу или нулю.\n" +
		"Результатом работы с римскими числами могут быть только положительные числа.\n" +
		"---------------------------------------------------------------------------- \n" +
		"Введите выражение (например, 3 + 4 или V - II):")

	// Используем bufio.NewReader для чтения всей строки целиком
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Убираем символ новой строки и лишние пробелы
	input = strings.TrimSpace(input)

	// Делим выражение на 2 операнда и один оператор
	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Print(len(parts))
		panic("Неверный формат ввода. Введите выражение вида a + b.")
	}
	num1Str, operator, num2Str := parts[0], parts[1], parts[2]
	// Преобразуем строки в целые числа
	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	// Проверяем, что оба значения успешно преобразованы в int
	if err1 == nil && err2 == nil {

		// Проверка на диапазон чисел от 1 до 10, т.к мы используем не натуральные, а целые числа, то отрицательный диапазон также захватываем
		if !((num1 >= 1 && num1 <= 10) || (num1 <= -1 && num1 >= -10) && (num2 >= 1 && num2 <= 10) || (num2 <= -1 && num2 >= -10)) {
			panic("Числа должны быть в диапазоне от 1 до 10 или от -1 до -10.")
		}

		result := calculate(num1, num2, operator)
		if result > 0 {
			panic("Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль.")
		}
		fmt.Println("Результат =", result)
	} else if isRomanNum(num1Str) && isRomanNum(num2Str) {
		num1 = romanToArabic[num1Str]
		num2 = romanToArabic[num2Str]

		// Проверка на диапазон римских чисел
		if !((num1 >= 1 && num1 <= 10) && (num2 >= 1 && num2 <= 10)) {
			panic("Римские числа должны быть в диапазоне от I до X.")
		}

		result := calculate(num1, num2, operator)
		if result <= 0 {
			panic("Результат римских чисел должен быть положительным.")
		}
		fmt.Println("Результат =", ArabicToRoman(result))
	} else {
		// Если операнды смешаны или имеют некорректный формат
		panic("Числа должны быть либо только арабскими, либо только римскими и при этом целые в диапазоне от 1 до 10 включительно. Повторите ввод.")
	}
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль невозможно.")
		}
		return num1 / num2
	default:
		panic("Неизвестный оператор.")
	}
}

// Проверка

func isRomanNum(s string) bool {
	_, exists := romanToArabic[s]
	return exists
}

func ArabicToRoman(s int) string {
	val := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	for i := 0; i < len(val); i++ {
		for s >= val[i] {
			s -= val[i]
			roman += syb[i]
		}
	}

	return roman

}
