package graph

import (
	"strings"
	"fmt"
	"sort"
)

type node struct {
	Val int
	weight *int
	Next *node
}

func (n *node) Weight() int {
	return *n.weight
}

// a graph that uses an adjacency list
type graph struct {
	Nodes map[int]*node
	Directed bool
	weighted bool
}

func New(directed bool) *graph {
	return &graph{
		Directed: directed,
		Nodes: map[int]*node{},
	}
}

func (g *graph) Insert(x, y int) {
	insertNode(g, x, y, nil, g.Directed)
}

func (g *graph) InsertWeighted(x, y, weight int) {
	insertNode(g, x, y, &weight, g.Directed)
}

func insertNode(g *graph, x, y int, w *int, directed bool) {
	newNode := &node{Val: y, weight: w}
	g.Nodes[y] = g.Nodes[y]
	newNode.Next = g.Nodes[x]
	g.Nodes[x] = newNode

	if !directed {
		insertNode(g, y, x, w, true)
	}
}

// String returns an adjacency list representation of a grpah
func (g *graph) String() string {
	b := &strings.Builder{}
	nodeKeys := func() []int {
		keys := []int{}
		for k, _ := range g.Nodes {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		return keys
	}()

	for _, k := range nodeKeys {
		fmt.Fprintf(b, "%d => ", k)
		for node := g.Nodes[k]; node != nil; node = node.Next {
			fmt.Fprintf(b, "%d -> ", node.Val)
		}
		fmt.Fprint(b, "\n")
	}
	return b.String()
}