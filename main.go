package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements создаёт слайс из size случайных положительных чисел.
// При нулевом или отрицательном размере возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(SIZE) + 1
	}
	return data
}

// maximum последовательно ищет наибольший элемент слайса.
// Для пустого слайса возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

// maxChunks делит слайс на восемь частей и ищет максимум каждой части
// в отдельной горутине. Короткие слайсы обрабатывает последовательно.
func maxChunks(data []int) int {
	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}

	maxima := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := range maxima {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			// Последняя часть включает остаток от деления длины слайса на CHUNKS.
			end = len(data)
		}

		go func(index int, chunk []int) {
			defer wg.Done()
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	// После завершения всех горутин в maxima записаны максимумы всех частей.
	wg.Wait()
	return maximum(maxima)
}

// main сравнивает время двух способов поиска на одном слайсе.
// Время поиска выводится в микросекундах.
func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	largest := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", largest, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	largest = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", largest, elapsed)
}
