package onehundreddays

import "fmt"
import "strconv"

func itoa(n int) string {
	return strconv.Itoa(n)
}

func equateArrays(xs, ys []string) bool {
	if xs == nil || ys == nil {
		return xs == nil && ys == nil
	}

	if len(xs) != len(ys) {
		return false
	}

	for i := range xs {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}

func reverseArray(xs []string) []string {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}

	return xs
}

func showArray(xs []string) string {
	return fmt.Sprintf("%v", xs)
}

func stringToUnitStrings(s string) []string {
	a := make([]string, len(s))

	for i, r := range s {
		a[i] = string(r)
	}

	return a
}

func to(start int, end int) []int {
	a := make([]int, end-start)

	for i := start; i < end; i++ {
		a[i-start] = i
	}

	return a
}

func fill(xs []int, x int) []int {
	for i := range xs {
		xs[i] = x
	}

	return xs
}

func sum(xs []int) int {
	total := 0

	for _, x := range xs {
		total += x
	}

	return total
}
