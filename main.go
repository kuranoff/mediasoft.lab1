package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	model := make([]int, 12)
	res := make([]int, 12)

	model = randomGenerator(len(model))
	fmt.Println(model)

	res = modelSort(model, len(model))
	fmt.Println(res)

}

/* Функция генерирования дискретной случайной величины */
func randomGenerator(lineup int) []int {
	numline := make([]int, lineup)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lineup; i++ {
		numline[i] = rand.Intn(100)
	}
	return numline
}

/* Функция пузырьковой сортировки */
func modelSort(num []int, size int) []int {
	for i := 0; i < size-1; i++ {
		for j := size - 1; j > i; j-- {
			if num[j-1] > num[j] {
				e := num[j-1]
				num[j-1] = num[j]
				num[j] = e
			}
		}
	}
	return num
}
