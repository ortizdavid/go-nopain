package units

import (
	"fmt"
	"testing"
)

func Test_LenghtUnits(t *testing.T) {
	fmt.Println(1 * METER) //1 m
	fmt.Println(5 * KILO_METER) //5 Km
	fmt.Println(0.7 * HECTO_METER) //0.7Hm
	fmt.Println(PETA_METER)
	fmt.Println(DECI_METER) // 1 dm
	fmt.Println(50 * CENTI_METER) // 50 cm
	fmt.Println(100 * MILI_METER) // 100 mm
	fmt.Println(0.5 * MICRO_METER) // 0.5 um
	fmt.Println(10 * NANO_METER) // 10 nm
	fmt.Println(10 * FEMTO_METER)
}
