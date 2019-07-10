package onehundreddays

import "fmt"
import "testing"

type mat struct {
	cost  int
	order string
	rows  int
	cols  int
}

func multiplications(chain []mat) mat {
	n := len(chain)
	aux := make([][]mat, n) // TODO: try using a map instead? map[string]int

	for i := range to(0, n) {
		aux[i] = make([]mat, n)
		aux[i][i] = chain[i]
	}

	// length of subchain
	for i := range to(1, n) {

		// starting index of subchain
		for j := range to(0, n-i) {
			best := 999999999

			// splitting point of subchain
			for k := range to(j, j+i) {

				// multiply subchains at splitting point
				l := aux[j][k]
				r := aux[k+1][j+i]
				cost := l.cost + r.cost + l.rows*l.cols*r.cols
				order := "(" + l.order + r.order + ")"

				if cost < best {
					best = cost
					aux[j][j+i] = mat{cost, order, l.rows, r.cols}
				}
			}
		}
	}

	return aux[0][n-1]
}

func testChain(t *testing.T, expectedOrder string, expectedCost int, chain []mat) {
	actual := multiplications(chain)
	if expectedOrder != actual.order || expectedCost != actual.cost {
		t.Errorf("order " + actual.order + " with cost " + itoa(actual.cost) + " does not match expected order " + expectedOrder + " with cost " + itoa(expectedCost))
	}
}

func Test002(t *testing.T) {
	// TODO: returning zero
	// testChain(t, "((AB)C)", 18000, []mat{
	// 	mat{0, "A", 10, 20},
	// 	mat{0, "B", 20, 30},
	// 	mat{0, "C", 30, 40},
	// })
}
