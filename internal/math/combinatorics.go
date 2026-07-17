package mathutil

// GCD returns the greatest common divisor of a and b via the Euclidean
// algorithm. GCD(0, b) = b and GCD(a, 0) = a, matching convention (0 is
// divisible by everything).
func GCD(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// LCM returns the least common multiple of a and b. Divides before
// multiplying to reduce overflow risk on large inputs; the division is exact
// since GCD(a, b) always divides a.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / GCD(a, b) * b
}

// LCMTable precomputes lcm(i, j) for all i, j in [1, mx]. O(mx^2) time and
// space -- only worth building when you need repeated O(1) lookups inside an
// mx-by-mx loop (e.g. a DP indexed by pairs of divisors). For occasional
// lookups, just call LCM directly.
func LCMTable(mx int) [][]int {
	t := make([][]int, mx+1)
	for i := range t {
		t[i] = make([]int, mx+1)
	}
	for i := 1; i <= mx; i++ {
		for j := 1; j <= mx; j++ {
			t[i][j] = LCM(i, j)
		}
	}
	return t
}

// PowTable precomputes base^i mod m for i in [0, mx] via iterative
// multiplication: O(mx) total, vs O(mx log mx) for mx independent calls to a
// fast-exponentiation routine. Use when you need base^i mod m for many
// values of i up to mx.
func PowTable(base, mx, mod int) []int {
	pow := make([]int, mx+1)
	pow[0] = 1 % mod
	for i := 1; i <= mx; i++ {
		pow[i] = pow[i-1] * base % mod
	}
	return pow
}

// PowMod computes base^exp mod m via binary exponentiation, O(log exp).
// Prefer over PowTable when you only need one (or a few) large, sparse
// exponents rather than a dense table.
func PowMod(base, exp, mod int) int {
	result := 1 % mod
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}