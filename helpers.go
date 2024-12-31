package helpers

import (
	"math/rand"
)

// Create slice of ints in the range from start to end, ordered from smallest to largest.
func CreateRange(start, end int) []int {
	size := end - start + 1
	rangeSlice := make([]int, size)
	for i := 0; i < size; i++ {
		rangeSlice[i] = start + i
	}
	return rangeSlice
}

// Shuffles the arrangement of ints in a slice.
func ShuffleSlice(slice []int) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
