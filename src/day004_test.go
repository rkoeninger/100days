package onehundreddays

import "strconv"
import "testing"

func countOneBits(value int) int {
	n := 0

	for value != 0 {
		value &= value - 1
		n++
	}

	return n
}

func testCountOneBits(t *testing.T, value, count int) {
	if count != countOneBits(value) {
		t.Errorf(strconv.Itoa(value) + " does not have " + strconv.Itoa(count) + " 1-bits")
	}
}

func Test004(t *testing.T) {
	testCountOneBits(t, 99, 4)
	testCountOneBits(t, 37, 3)
	testCountOneBits(t, 255, 8)
	testCountOneBits(t, -1, 64)
}
