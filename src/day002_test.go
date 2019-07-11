package onehundreddays

import "math"
import "testing"

type mat struct {
	cost  int
	order string
	rows  int
	cols  int
}

func multiplications(chain []mat) mat {
	n := len(chain)
	aux := make([][]mat, n)

	for _, i := range to(0, n) {
		aux[i] = make([]mat, i+1)
		aux[i][i] = chain[i]
	}

	// length of subchain
	for _, i := range to(1, n) {

		// starting index of subchain
		for _, j := range to(0, n-i) {
			best := math.MaxInt32

			// splitting point of subchain
			for _, k := range to(j, j+i) {

				// multiply subchains at splitting point
				l := aux[k][j]
				r := aux[j+i][k+1]
				cost := l.cost + r.cost + l.rows*l.cols*r.cols
				order := "(" + l.order + r.order + ")"

				if cost < best {
					best = cost
					aux[j+i][j] = mat{cost, order, l.rows, r.cols}
				}
			}
		}
	}

	return aux[n-1][0]
}

func testChain(t *testing.T, expected mat, chain []mat) {
	actual := multiplications(chain)
	if expected.order != actual.order || expected.cost != actual.cost {
		t.Errorf("order " + actual.order + " with cost " + itoa(actual.cost) + " does not match expected order " + expected.order + " with cost " + itoa(expected.cost))
	}
}

func Test002(t *testing.T) {
	testChain(t, mat{18000, "((AB)C)", 10, 40}, []mat{
		mat{0, "A", 10, 20},
		mat{0, "B", 20, 30},
		mat{0, "C", 30, 40},
	})
}
