package random

import (
	"math/rand"
)


// Int generates a random integer between min (inclusive) and max (exclusive).
func Int(min, max int) int {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return rand.Intn(max-min) + min
}

// Int32 generates a random 32-bit integer between min (inclusive) and max (exclusive).
func Int32(min, max int32) int32 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return rand.Int31n(max-min) + min
}

// Int64 generates a random 64-bit integer between min (inclusive) and max (exclusive).
func Int64(min, max int64) int64 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return rand.Int63n(max-min) + min
}

// Float32 generates a random float32 between min (inclusive) and max (exclusive).
func Float32(min, max float32) float32 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return min + rand.Float32()*(max-min)
}

// Float64 generates a random float64 between min (inclusive) and max (exclusive).
func Float64(min, max float64) float64 {
	if min >= max {
		panic("Invalid range. min should be less than max.")
	}
	seedRandom()
	return min + rand.Float64()*(max-min)
}
