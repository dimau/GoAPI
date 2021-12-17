package utils

import "testing"

// Функция используется для инициализации среза целых чисел заданным числом чисел
func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++

	}
	return result
}

// Бенчмарк тест-кейс для функции Sort10
func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
