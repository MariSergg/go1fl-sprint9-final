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

// generateRandomElements создаёт слайс случайных положительных целых чисел.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

// maximum возвращает максимальное значение из слайса.
// Для пустого слайса возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]

	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks находит максимум с помощью восьми горутин.
func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)

		go func(index int, chunk []int) {
			defer wg.Done()

			chunkMaximums[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение последовательно")

	start := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение: %d\nВремя поиска: %d мкс\n",
		maxSingle,
		elapsedSingle,
	)

	fmt.Printf(
		"Ищем максимальное значение в %d горутинах\n",
		CHUNKS,
	)

	start = time.Now()
	maxParallel := maxChunks(data)
	elapsedParallel := time.Since(start).Microseconds()

	fmt.Printf(
		"Максимальное значение: %d\nВремя поиска: %d мкс\n",
		maxParallel,
		elapsedParallel,
	)

	if maxSingle != maxParallel {
		fmt.Println("Ошибка: результаты поиска не совпадают")
	}
}
