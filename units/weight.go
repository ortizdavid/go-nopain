package units

// Constants for various mass units, all based on the gram.
const (
	GRAM       float64 = 1
	MILI_GRAM  float64 = GRAM / 1000
	HECTO_GRAM float64 = GRAM * 100
	KILO_GRAM  float64 = GRAM * 1000
	TONNE      float64 = KILO_GRAM * 1000

	OUNCE      float64 = GRAM / 28.3495   // International avoirdupois ounce
	POUND      float64 = KILO_GRAM / 0.453592   // International avoirdupois pound
)
