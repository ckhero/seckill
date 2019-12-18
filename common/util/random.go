package util

import (
	"math/rand"
	"time"
)

func RandRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}