package units

const (
	METER 			float64 = 1
	DECA_METER 		float64 = METER * 10
	HECTO_METER 	float64	= METER * 100
	KILO_METER 		float64 = METER * 1000
	PENTA_METER 	float64 = METER * 1_000_000_000_000_000

	DECI_METER 		float64 = METER / 10
	CENTI_METER 	float64 = METER / 100
	MILI_METER 		float64 = METER / 1000
	MICRO_METER		float64 = METER / 1_000_000
	NANO_METER		float64 = METER / 1_000_000_000
	FENTO_METER		float64 = METER / 1_000_000_000_000_000

	INCH 			float64 = CENTI_METER / 2.54
	FOOT 			float64 = CENTI_METER / 30.48
	YARD			float64 = METER / 0.914
	MILES 			float64 = KILO_METER / 1.60934
)
