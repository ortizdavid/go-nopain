package units

import (
	"fmt"
	"testing"
)

func Test_CapacityUnits(t *testing.T) {
	fmt.Println(1 * LITER) // 1 l
	fmt.Println(DECI_LITER)
	fmt.Println(33 * CENTI_LITER) // 33 cl
	fmt.Println(100 * MILI_LITER) // 100 ml
	fmt.Println(DECA_LITER)
	fmt.Println(HECTO_LITER)
	fmt.Println(KILO_LITER)
}
