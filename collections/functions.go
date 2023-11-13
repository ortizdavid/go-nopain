package collections

func ContainsString(stringSlice []string, str string) bool {
	for _, item := range stringSlice {
		if str == item {
			return true
		}
	}
	return false
}