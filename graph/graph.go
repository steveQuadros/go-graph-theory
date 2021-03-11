package graph

import (
	"strings"
	"fmt"
)

type Node struct {
	Val int
	weight *int
	Next *Node
}

func (n *Node) Weight() int {
	return *n.weight
}

// a graph that uses an adjacency list
type Graph struct {
	Nodes []*Node
	Directed bool
	Weighted bool
}

type Options struct {
	Directed, Weighted bool
}

// create a new graph with nvertices 0 indexed
func New(nvertices int, o *Options) *Graph {
	return &Graph{
		Directed: o.Directed,
		Weighted: o.Weighted,
		Nodes: make([]*Node, nvertices, nvertices),
	}
}

func (g *Graph) Insert(x, y int) {
	insertNode(g, x, y, nil, g.Directed)
}

func (g *Graph) InsertWeighted(x, y, weight int) {
	insertNode(g, x, y, &weight, g.Directed)
}

func insertNode(g *Graph, x, y int, w *int, directed bool) {
	newNode := &Node{Val: y, weight: w}
	newNode.Next = g.Nodes[x]
	g.Nodes[x] = newNode

	if !directed {
		insertNode(g, y, x, w, true)
	}
}

// String returns an adjacency list representation of a grpah
func (g *Graph) String() string {
	b := &strings.Builder{}

	for i, node := range g.Nodes {
		fmt.Fprintf(b, "%d => ", i)
		for node != nil {
			fmt.Fprintf(b, "%d -> ", node.Val)
			node = node.Next
		}
		fmt.Fprint(b, "\n")
	}
	return b.String()
}