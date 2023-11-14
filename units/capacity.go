package units

const (
	LITER 		 	float64 = 1
	DECI_LITER   	float64 = LITER / 10
	CENTI_LITER		float64 = LITER / 100
	MILI_LITER 		float64 = LITER / 1000

	DECA_LITER 		float64 = LITER * 100
	HECTO_LITER 	float64 = LITER * 100
	KILO_LITER 		float64 = LITER * 1000

	FLUID_OUNCE  float64 = 0.0295735 * LITER      
    CUP          float64 = 0.236588 * LITER       
    PINT         float64 = 0.473176 * LITER        
    QUART        float64 = 0.946353 * LITER      
    GALLON       float64 = 3.78541 * LITER
)