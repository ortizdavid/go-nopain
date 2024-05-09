package units

import (
	"fmt"
	"testing"
)

func Test_StorageUnits(t *testing.T) {
	fmt.Println(8 * BIT) // 8 bits
	fmt.Println(4 * BYTE) // 4 B
	fmt.Println(45 * KILO_BYTE) // 45 KB
	fmt.Println(20 * MEGA_BYTE)
	fmt.Println(3 * GIGA_BYTE) // 3 GB
	fmt.Println(9 * TERA_BYTE) // 9 TB
	fmt.Println(PETA_BYTE)
	fmt.Println(EXA_BYTE)
	fmt.Println(ZETTA_BYTE)
	fmt.Println(YOTTA_BYTE)
}
