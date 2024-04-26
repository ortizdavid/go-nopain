package units

// Constants for various data storage units, all based on the byte.
const (
	BIT        float64 = BYTE / 8
	BYTE       float64 = 1
	KILO_BYTE  float64 = BYTE * 1024
	MEGA_BYTE  float64 = KILO_BYTE * 1024
	GIGA_BYTE  float64 = MEGA_BYTE * 1024
	TERA_BYTE  float64 = GIGA_BYTE * 1024
	PETA_BYTE  float64 = TERA_BYTE * 1024
	EXA_BYTE   float64 = PETA_BYTE * 1024
	ZETTA_BYTE float64 = EXA_BYTE * 1024
	YOTTA_BYTE float64 = ZETTA_BYTE * 1024
)
