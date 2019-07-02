package utils

func MinInt(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func MaxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
