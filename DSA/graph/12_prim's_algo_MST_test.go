package graph

import "testing"

func TestPrimsAlgoMST(t *testing.T) {

	edges := [][]int{{1, 2, 2}, {1, 4, 6}, {2, 1, 2}, {2, 3, 3}, {2, 4, 8}, {2, 5, 5}, {3, 2, 3},
		{3, 5, 7}, {4, 1, 6}, {4, 2, 8}, {4, 5, 9}, {5, 2, 5}, {5, 3, 7}, {5, 4, 9}}
	primsAlgoMST(5, edges)

}
