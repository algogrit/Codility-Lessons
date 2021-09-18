package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneNeighbors(lookup map[int]*Node, node *Node) *Node {
	for idx, neighbor := range node.Neighbors {
		clonedNeighbor, ok := lookup[neighbor.Val]
		if !ok {
			clonedNeighbor = new(Node)
			clonedNeighbor.Val = neighbor.Val
			clonedNeighbor.Neighbors = make([]*Node, len(neighbor.Neighbors))
			copy(clonedNeighbor.Neighbors, neighbor.Neighbors)

			lookup[clonedNeighbor.Val] = clonedNeighbor
			clonedNeighbor = cloneNeighbors(lookup, clonedNeighbor)
		}
		node.Neighbors[idx] = clonedNeighbor
	}
	return node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	clonedNode := new(Node)
	clonedNode.Val = node.Val
	clonedNode.Neighbors = make([]*Node, len(node.Neighbors))
	copy(clonedNode.Neighbors, node.Neighbors)

	return cloneNeighbors(map[int]*Node{clonedNode.Val: clonedNode}, clonedNode)
}

func main() {
	// An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.
	// adjList = [[2,4],[1,3],[2,4],[1,3]]
	node1 := &Node{Val: 1, Neighbors: make([]*Node, 2)}
	node2 := &Node{Val: 2, Neighbors: make([]*Node, 2)}
	node3 := &Node{Val: 3, Neighbors: make([]*Node, 2)}
	node4 := &Node{Val: 4, Neighbors: make([]*Node, 2)}

	node1.Neighbors[0] = node2
	node1.Neighbors[1] = node4

	node2.Neighbors[0] = node1
	node2.Neighbors[1] = node3

	node3.Neighbors[0] = node2
	node3.Neighbors[1] = node4

	node4.Neighbors[0] = node1
	node4.Neighbors[1] = node3

	fmt.Printf("%p %#v\n", node1, node1)

	newGraph := cloneGraph(node1)

	fmt.Printf("%p %#v\n", newGraph, newGraph)
}
