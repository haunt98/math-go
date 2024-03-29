package math

// GCD is greatest common divisor
func GCD(a, b int64) int64 {
	// Odd cases

	if a == 0 && b == 0 {
		return 0
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	// Only deal with positive numbers
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	// Euler algorithm

	if a < b {
		a, b = b, a
	}

	for b != 0 {
		a = a % b
		a, b = b, a
	}

	return a
}
