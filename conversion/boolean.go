package conversion

import "strconv"

func StringToBool(strBool string) (bool, error) {
	return strconv.ParseBool(strBool)
}

func BoolToString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}