package graph

import "testing"

func TestBellmanFordShortestPath(t *testing.T) {

	edges := [][]int{{1, 2, 2}, {1, 3, 2}, {2, 3, -1}}
	bellmanFordShortestPath(3, edges, 1, 3)

	edges = [][]int{{1, 13, 25}, {2, 1, 41}, {2, 8, -69}, {3, 11, 113}, {5, 3, 3}, {5, 10, 107}, {5, 12, 121}, {6, 5, 56},
		{6, 9, 125}, {8, 10, 90}, {9, 5, 86}, {9, 7, 98}, {10, 11, 97}, {11, 10, 69}, {11, 10, 139},
		{11, 12, 50}, {12, 3, 25}, {12, 14, 25}, {14, 9, 13}}
	bellmanFordShortestPath(14, edges, 1, 13)

}
