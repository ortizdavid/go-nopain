package units

// Constants for various volume units, all based on the liter.
const (
	LITER         float64 = 1
	DECI_LITER    float64 = LITER / 10
	CENTI_LITER   float64 = LITER / 100
	MILI_LITER    float64 = LITER / 1000

	DECA_LITER    float64 = LITER * 10
	HECTO_LITER   float64 = LITER * 100
	KILO_LITER    float64 = LITER * 1000

	FLUID_OUNCE   float64 = 0.0295735 * LITER   // U.S. fluid ounce
	CUP           float64 = 0.236588 * LITER    // U.S. customary cup
	PINT          float64 = 0.473176 * LITER    // U.S. liquid pint
	QUART         float64 = 0.946353 * LITER    // U.S. liquid quart
	GALLON        float64 = 3.78541 * LITER     // U.S. liquid gallon
)
