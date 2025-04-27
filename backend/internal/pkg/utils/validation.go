package utils

func RoundToCents(value float64) float64 {
	return float64(int(value*100+0.5)) / 100
}

func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !(r >= 'a' && r <= 'z') &&
			!(r >= 'A' && r <= 'Z') &&
			!(r >= '0' && r <= '9') &&
			!(r == ' ') {
			return false
		}
	}
	return true
}
