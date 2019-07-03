package one_hundred_days

import "testing"

func hanoi_recur(height int, from, to, through string) []string {
	if height > 0 {
		return append(
			hanoi_recur(height-1, from, through, to),
			append(
				[]string{from + " => " + to},
				hanoi_recur(height-1, through, to, from)...)...)
	}
	return []string{}
}

func hanoi(height int) []string {
	return hanoi_recur(height, "left", "right", "center")
}

func array_equal(xs, ys []string) bool {
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

func test_hanoi(t *testing.T, height int, expected []string) {
	if !array_equal(expected, hanoi(height)) {
		t.FailNow()
	}
}

func Test001(t *testing.T) {
	test_hanoi(t, 1, []string{
		"left => right",
	})

	test_hanoi(t, 2, []string{
		"left => center",
		"left => right",
		"center => right",
	})

	test_hanoi(t, 3, []string{
		"left => right",
		"left => center",
		"right => center",
		"left => right",
		"center => left",
		"center => right",
		"left => right",
	})

	test_hanoi(t, 4, []string{
		"left => center",
		"left => right",
		"center => right",
		"left => center",
		"right => left",
		"right => center",
		"left => center",
		"left => right",
		"center => right",
		"center => left",
		"right => left",
		"center => right",
		"left => center",
		"left => right",
		"center => right",
	})
}
