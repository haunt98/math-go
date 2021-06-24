package math

// LCM is least common multiple
func LCM(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}

	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	return a * b / GCD(a, b)
}
