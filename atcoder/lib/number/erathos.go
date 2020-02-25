package number

import "math"

func Erathos(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		isPrime[i] = true
	}
	upperBound := int(math.Sqrt(float64(n + 1)))
	for i := 2; i < upperBound; i++ {
		if !isPrime[i] {
			continue
		}
		d := i + i
		for d < n+1 {
			isPrime[d] = false
			d += i
		}
	}
	return isPrime
}
