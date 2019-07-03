package one_hundred_days

import "strconv"
import "testing"

func count_one_bits(value int) int {
	n := 0
	for value != 0 {
		value &= value - 1
		n++
	}
	return n
}

func test_count_one_bits(t *testing.T, value, count int) {
	if count != count_one_bits(value) {
		t.Errorf(strconv.Itoa(value) + " does not have " + strconv.Itoa(count) + " 1-bits")
		t.FailNow()
	}
}

func Test004(t *testing.T) {
	test_count_one_bits(t, 99, 4)
	test_count_one_bits(t, 37, 3)
	test_count_one_bits(t, 255, 8)
	test_count_one_bits(t, -1, 64)
}
