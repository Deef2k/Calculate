package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	expression := "hello" + "world" // Пример выражения: "hello" + "world"
	result := evaluateExpression(expression)
	fmt.Println("Результат:", result)
}

func evaluateExpression(expression string) string {
	// Разбиваем выражение на операнды и оператор
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		panic("Неверное выражение")
	}

	operand1 := parts[0]
	operator := parts[1]
	operand2 := parts[2]

	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return subtractStrings(operand1, operand2)
	case "*":
		return multiplyString(operand1, operand2)
	case "/":
		return divideString(operand1, operand2)
	default:
		panic("Неподдерживаемая операция")
	}
}

func subtractStrings(a, b string) string {
	// Удаляем подстроку b из строки a
	return strings.ReplaceAll(a, b, "")
}

func multiplyString(s string, n string) string {
	// Повторяем строку s n раз
	count := parseInt(n)
	return strings.Repeat(s, count)
}

func divideString(s string, n string) string {
	// Деление строки s на число n
	count := parseInt(n)
	length := len(s)
	newLength := int(math.Ceil(float64(length) / float64(count)))
	if newLength > 40 {
		return s[:40] + "..."
	}
	return s[:newLength]
}

func parseInt(s string) int {
	// Преобразование строки в целое число
	num, err := strconv.Atoi(s)
	if err != nil {
		panic("Неверное число")
	}
	if num < 1 || num > 10 {
		panic("Число должно быть от 1 до 10")
	}
	return num
}
