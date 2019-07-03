package count1bits_test

import "testing"

func count1bits(value int) int {
	n := 0
	for value > 0 {
		value &= value - 1
		n++
	}
	return n
}

func Test(t *testing.T) {
	if 4 != count1bits(99) {
		t.FailNow()
	}
}