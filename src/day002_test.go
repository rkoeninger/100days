package onehundreddays

import "testing"

type mat struct {
	Name string
	Rows int
	Cols int
}

type multResult struct {
	Cost  int
	Order string
	Rows  int
	Cols  int
}

func multiplications(chain []mat) multResult {
	n := len(chain)

	/*
	  # single matrix chain has zero cost
	  aux = {(i, i): (0,) + chain[i] for i in range(n)}

	  # i: length of subchain
	  for i in range(1, n):

	    # j: starting index of subchain
	    for j in range(0, n - i):
	      best = float('inf')

	      # k: splitting point of subchain
	      for k in range(j, j + i):
	        # multiply subchains at splitting point
	        lcost, lname, lrow, lcol = aux[j, k]
	        rcost, rname, rrow, rcol = aux[k + 1, j + i]
	        cost = lcost + rcost + lrow * lcol * rcol
	        var = '(%s%s)' % (lname, rname)

	        # pick the best one
	        if cost < best:
	          best = cost
	          aux[j, j + i] = cost, var, lrow, rcol

	  return dict(zip(['cost', 'order', 'rows', 'cols'], aux[0, n - 1]))
	*/

	return multResult{n, "", 0, 0}
}

func testChain(chain []mat, expected multResult) {

}

func Test002(t *testing.T) {

}
