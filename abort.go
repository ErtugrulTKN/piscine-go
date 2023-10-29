package piscine

func Abort(a, b, c, d, e int) int { // Sorting yaparak 5 argümandan ortadakini al.
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	if a > d {
		a, d = d, a
	}
	if b > d {
		b, d = d, b
	}
	if c > d {
		c, d = d, c
	}
	if a > e {
		a, e = e, a
	}
	if b > e {
		b, e = e, b
	}
	if c > e {
		c, e = e, c
	}
	if d > e {
		d, e = e, d
	}

	// The median
	return c
}
