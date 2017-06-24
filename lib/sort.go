package lib

import (
	"math/rand"
)

func GenRandomFloatSlice(n int) []float32 {
	randSlice := make([]float32, 0)
	for i := 0; i < n; i++ {
		randSlice = append(randSlice, rand.Float32())
	}
	return randSlice
}
