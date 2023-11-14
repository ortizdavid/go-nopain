
package random

import (
	"math/rand"
	"time"
)

func seedRandom() {
	rand.Seed(time.Now().UnixNano())
}

func Int(min, max int) int {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return rand.Intn(max-min) + min
}

func Int32(min, max int32) int32 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return rand.Int31n(max-min) + min
}

func Int64(min, max int64) int64 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return rand.Int63n(max-min) + min
}

func Float32(min, max float32) float32 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return min + rand.Float32()*(max-min)
}

func Float64(min, max float64) float64 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return min + rand.Float64()*(max-min)
}
