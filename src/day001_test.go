package onehundreddays

import "testing"

func hanoiRecur(height int, from, to, through string) []string {
	if height > 0 {
		return append(
			hanoiRecur(height-1, from, through, to),
			append(
				[]string{from + " => " + to},
				hanoiRecur(height-1, through, to, from)...)...)
	}

	return []string{}
}

func hanoi(height int) []string {
	return hanoiRecur(height, "left", "right", "center")
}

func testHanoi(t *testing.T, height int, expected []string) {
	actual := hanoi(height)
	if !equateArrays(expected, actual) {
		t.Errorf("move sequence " + showArray(actual) + " does not match expected " + showArray(expected))
	}
}

func Test001(t *testing.T) {
	testHanoi(t, 1, []string{
		"left => right",
	})

	testHanoi(t, 2, []string{
		"left => center",
		"left => right",
		"center => right",
	})

	testHanoi(t, 3, []string{
		"left => right",
		"left => center",
		"right => center",
		"left => right",
		"center => left",
		"center => right",
		"left => right",
	})

	testHanoi(t, 4, []string{
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
