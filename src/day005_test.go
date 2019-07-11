package onehundreddays

import "math"
import "testing"

func eratosthenes(n int) int {
	n = (n + 1) >> 1
	i, j := 1, 3
	p := fill(make([]int, n), 1)

	for j < int(math.Ceil(math.Sqrt(float64(2 * n)))) {
		if p[i] > 0 {
			for k := j * j >> 1; k < n; k += j {
				p[k] = 0
			}
		}
		i += 1
		j += 2
	}

	return sum(p)
}

func testEratosthenes(t *testing.T, n, expected int) {
	actual := eratosthenes(n)
	if actual != expected {
		t.Errorf("there are " + itoa(expected) + " prime numbers under " + itoa(n) + ", not " + itoa(actual))
	}
}

// https://primes.utm.edu/howmany.html
func Test005(t *testing.T) {
	testEratosthenes(t, 10, 4)
	testEratosthenes(t, 100, 25)
	testEratosthenes(t, 1000, 168)
	testEratosthenes(t, 10000, 1229)
	testEratosthenes(t, 100000, 9592)
	testEratosthenes(t, 1000000, 78498)
}
