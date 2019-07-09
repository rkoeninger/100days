package onehundreddays

import "strconv"
import "testing"

func findPivot(a []string) int {
	if len(a) > 1 {
		for i := len(a) - 2; i >= 0; i-- {
			if a[i] < a[i+1] {
				return i
			}
		}
	}

	return -1
}

func nextPermutation(a []string) []string {
	i := findPivot(a)

	if i == -1 {
		return reverseArray(a)
	}

	for j := len(a) - 1; j >= i; j-- {
		if a[i] < a[j] {
			a[i], a[j] = a[j], a[i]
			return append(a[:i+1], reverseArray(a[i+1:])...)
		}
	}

	return a
}

func testPermutation(t *testing.T, input, expected string) {
	inputArray := stringToUnitStrings(input)
	expectedArray := stringToUnitStrings(expected)
	actualArray := nextPermutation(inputArray)
	if !equateArrays(actualArray, expectedArray) {
		t.Errorf("permutation " + showArray(actualArray) + " is not the expected " + showArray(expectedArray))
		t.Errorf("    with pivot i = " + strconv.Itoa(findPivot(inputArray)))
	}
}

func Test003(t *testing.T) {
	testPermutation(t, "FADE", "FAED")
	testPermutation(t, "PHONE", "PNEHO")
}
