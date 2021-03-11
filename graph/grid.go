package graph

import (
	"strings"
	"fmt"
)

type Grid struct {
	NVertices int
	Nodes [][]bool
	Directed bool
}

// Create a matrix representation of a graph with vertices [0, n) (0 -> n-1)
func NewGrid(nvertices int, o *Options) *Grid {
	nodes := make([][]bool, nvertices, nvertices)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = make([]bool, nvertices)
	}

	return &Grid{
		NVertices: nvertices,
		Directed: o.Directed,
		Nodes: nodes,
	}
}

func (g *Grid) Insert(x,y int) {
	g.Nodes[x][y] = true
	if !g.Directed {
		g.Nodes[y][x] = true
	}
}

func (g *Grid) String() string {
	b := &strings.Builder{}
	// print gap for alignment
	fmt.Fprintf(b, "   ")
	// print header with col nums
	for i := 0; i < len(g.Nodes); i++ {
		fmt.Fprintf(b, "%3d ", i)
	}
	fmt.Fprint(b, "\n")

	for i := 0; i < len(g.Nodes); i++ {
		fmt.Fprintf(b, "%-3d ", i)
		for j := 0; j < len(g.Nodes); j++ {
			var val int
			if g.Nodes[i][j] {
				val = 1
			} else {
				val = 0
			}
			fmt.Fprintf(b, "%3d ", val)
		}
		fmt.Fprint(b, "\n")
	}
	return b.String()
}