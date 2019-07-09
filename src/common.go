package onehundreddays

import "fmt"

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
