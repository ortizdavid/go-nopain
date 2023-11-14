package units

import (
	"fmt"
	"testing"
)


func TestWeigthUnits(t *testing.T) {
	fmt.Println(1 * GRAM) // 1 g
	fmt.Println(2 * MILI_GRAM) // 2 mg
	fmt.Println(HECTO_GRAM)
	fmt.Println(10 * KILO_GRAM) // 10 Kg
	fmt.Println(16 * TONNE) // 16 t
	fmt.Println(8 * POUND)
	fmt.Println(OUNCE)
}
