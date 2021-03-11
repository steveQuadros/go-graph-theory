package topsort

import (
	"github.com/stevequadros/go-graph-theory/graph"
)

// By Definition all rooted trees have a topological ordering
// most of the time in CS when discussing trees, we are discussing a rooted tree, and often times also a rooted out-tree (aborescence)

// Use Kahn's algorithm to topsort
// returns nil if there is a cycle
func TopsortKahn(g *graph.Graph) []int {
	nodes := func(nodes []*graph.Node) map[int]*graph.Node {
		m := map[int]*graph.Node{}
		for i, n := range nodes {
			m[i] = n
		}
		return m
	}(g.Nodes)

	inDegree := func(nodes []*graph.Node) []int {
		id := make([]int, len(nodes))
		for i := 0; i < len(id); i++ {
			for node := nodes[i]; node != nil; node = node.Next {
				id[node.Val]++
			}
		}
		return id
	}(g.Nodes)
	
	queue := []int{}
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := make([]int, len(nodes))
	idx := 0
	for len(queue) != 0 {
		at := queue[0]
		queue = queue[1:]

		order[idx] = at
		idx++

		for node := nodes[at]; node != nil; node = node.Next {
			inDegree[node.Val]--
			if inDegree[node.Val] == 0 {
				queue = append(queue, node.Val)
				
			}
		}
		delete(nodes, at)
	}

	// cycle detected since we didn't make it through all nodes
	if idx != len(g.Nodes) {
		return nil
	}

	return order
}

// TopSort will find the topological ordering of a Tree ()
func TopSortDFS(g *graph.Graph) []int {
	order := make([]int, len(g.Nodes))
	idx := len(g.Nodes)-1
	visited := make([]bool, len(g.Nodes))

	var dft func(int)
	dft = func(at int) {
		if visited[at] {
			return
		}
		visited[at] = true
		for node := g.Nodes[at]; node != nil; node = node.Next {
			dft(node.Val)
		}
		order[idx] = at
		idx--
		return
	}

	for i := 0; i < len(g.Nodes); i++ {
		if !visited[i] {
			dft(i)
		}
	}

	return order
}