package main

import (
	"fmt"
)

func main() {
	addEdge(0, 1, 0)
	addEdge(0, 2, 0)
	addEdge(1, 2, 0)
	fmt.Println(adj)

	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 5}, {3, 5}, {4, 6}}

	printAdjacency(edges)
}

var adj = make(map[int][]int)

/*
direction: 0 = undirected graph, 1 = directed graph
*/
func addEdge(u, v, direction int) {

	adj[u] = append(adj[u], v)

	if direction == 0 {
		adj[v] = append(adj[v], u)
	}
}

func printAdjacency(edges [][]int) {

	adj = make(map[int][]int)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	fmt.Println(adj)
}
