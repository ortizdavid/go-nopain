package random

import (
	"math/rand"
	"time"
)


// seedRandom seeds the random number generator with the current time.
func seedRandom() {
	rand.Seed(time.Now().UnixNano())
}
