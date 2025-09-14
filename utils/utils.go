package utils

import (
	"math/rand"
	"time"
)

func RandomValueBetween(x, y float32) float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	values := []float32{x, y} // -1.0 has double the chance
	return values[r.Intn(len(values))]
}
