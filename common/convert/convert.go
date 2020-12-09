package convert

//IntToBool transforma enteros  bool
func IntToBool(n int8) bool {
	if n > 0 || n < 0 {
		return true
	}
	return false
}

//BoolToInt transforma bool a enteros
func BoolToInt(b bool) int8 {
	if b {
		return 1
	}
	return 0
}
