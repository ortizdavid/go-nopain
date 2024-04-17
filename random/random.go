package random

import (
	"math/rand"
	"time"
)


func seedRandom() {
	rand.Seed(time.Now().UnixNano())
}