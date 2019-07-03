package one_hundred_days

import (
	"fmt"
	"testing"
)

func hanoi_recur(height int, from string, to string, through string) []string {
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

func show(moves []string) {
	for i := 0; i < len(moves); i++ {
		fmt.Println(moves[i])
	}
	fmt.Println()
}

func Test001(t *testing.T) {
	show(hanoi(1))
	show(hanoi(2))
	show(hanoi(3))
	show(hanoi(4))
}
