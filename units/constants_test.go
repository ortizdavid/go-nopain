package units

import (
	"fmt"
	"testing"
)

func TestStorageUnits(t *testing.T) {
	fmt.Println(BIT * 8)
	fmt.Println(BYTE / 8)
	fmt.Println(KILO_BYTE)
	fmt.Println(MEGA_BYTE)
	fmt.Println(GIGA_BYTE)
	fmt.Println(TERA_BYTE)
	fmt.Println(PETA_BYTE)
	fmt.Println(EXA_BYTE)
	fmt.Println(ZETA_BYTE)
	fmt.Println(YOTA_BYTE)
}

func TestLenghtUnits(t *testing.T) {
	fmt.Println(METER)
	fmt.Println(KILO_METER)
	fmt.Println(HECTO_METER)
	fmt.Println(PENTA_METER)
	fmt.Println(DECI_METER)
	fmt.Println(CENTI_METER)
	fmt.Println(MILI_METER)
	fmt.Println(MICRO_METER)
	fmt.Println(NANO_METER)
	fmt.Println(FENTO_METER)
}

func TestCapacityUnits(t *testing.T) {
	fmt.Println(LITER)
	fmt.Println(DECI_LITER)
	fmt.Println(CENTI_LITER)
	fmt.Println(MILI_LITER)
	fmt.Println(DECA_LITER)
	fmt.Println(HECTO_LITER)
	fmt.Println(KILO_LITER)
}

func TestWeigthUnits(t *testing.T) {

}
