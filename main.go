package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {

	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	src := rand.NewSource(time.Now().Unix())
	for j := 0; j < size; j++ {
		data[j] = int(src.Int63())
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	
	if len(data) == 0 {
		log.Println("len of slice is zero")
		return 0
	}

	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	
	lenSlice := len(data)

	if lenSlice < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup

	maxSlices := make([]int, CHUNKS)

	var i int

	for i = 0; i < (CHUNKS - 1); i++ {
		startIdx := lenSlice / CHUNKS * i
		endIdx := startIdx + (lenSlice / CHUNKS)

		partSlice := data[startIdx:endIdx]

		wg.Add(1)

		go func(partSlice []int, i int) {
			defer wg.Done()
			maxSlices[i] = maximum(partSlice)
		}(partSlice, i)
	}

	wg.Wait()
	idx := lenSlice / CHUNKS * (CHUNKS - 1)
	maxSlices[CHUNKS-1] = maximum(data[idx:])

	maxValue := maximum(maxSlices)
	return maxValue
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
