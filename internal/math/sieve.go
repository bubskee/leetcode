package math

// MobiusSieve computes the Möbius function μ(n) for all n in [0, mx] using a
// linear (Euler) sieve: O(mx) time, O(mx) space.
//
//	μ(1) = 1
//	μ(n) = 0        if n has a squared prime factor
//	μ(n) = (-1)^k   if n is squarefree with k distinct prime factors
//
// mu[0] is unused (left as 0); valid indices are 1..mx.
func MobiusSieve(mx int) []int {
	mu := make([]int, mx+1)
	isComposite := make([]bool, mx+1)
	primes := make([]int, 0)

	if mx >= 1 {
		mu[1] = 1
	}
	for i := 2; i <= mx; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
			mu[i] = -1
		}
		for _, p := range primes {
			if i*p > mx {
				break
			}
			isComposite[i*p] = true
			if i%p == 0 {
				mu[i*p] = 0
				break // i's smallest prime factor -- stop to keep this linear
			}
			mu[i*p] = -mu[i]
		}
	}
	return mu
}

// Primes returns all primes in [2, mx] via the same linear-sieve pass used by
// MobiusSieve. Kept as a separate function rather than a shared helper with a
// callback -- for class-notes purposes, seeing the bare loop twice is more
// useful for internalizing the mechanics than one DRY-but-abstracted version.
func Primes(mx int) []int {
	isComposite := make([]bool, mx+1)
	primes := make([]int, 0)
	for i := 2; i <= mx; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > mx {
				break
			}
			isComposite[i*p] = true
			if i%p == 0 {
				break
			}
		}
	}
	return primes
}

// DivisorSumTransform takes cnt where cnt[i] is a per-value quantity (e.g.
// count of elements equal to i) and transforms it in place so cnt[i] becomes
// the sum over all multiples of i (e.g. count of elements divisible by i).
// O(mx log mx) -- harmonic series over divisors.
//
// This is the building block Möbius inversion undoes: divisor-sum counts ->
// exact-value counts.
func DivisorSumTransform(cnt []int) {
	mx := len(cnt) - 1
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			cnt[i] += cnt[j]
		}
	}
}
