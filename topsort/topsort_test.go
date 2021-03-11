package topsort

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stevequadros/go-graph-theory/graph"
)

func TestTopSort(t *testing.T) {
	// note we have an unconnected node '6' with 0 inDegrees here as well
	g := graph.New(7, &graph.Options{Directed: true})
	g.Insert(0,1)
	g.Insert(0,2)
	g.Insert(0,3)
	g.Insert(1,4)
	g.Insert(1,5)
	g.Insert(3,5)
	g.Insert(5,4)

	assert.Equal(t, []int{0,6,3,2,1,5,4}, TopsortKahn(g))

	t.Run("with cycle returns nil", func(t *testing.T) {
		// create cycle
		g.Insert(4,5)

		assert.Equal(t, []int(nil), TopsortKahn(g))
	})
}

func TestTopSortDFS(t *testing.T) {
	g := graph.New(7, &graph.Options{Directed: true})
	g.Insert(0,1)
	g.Insert(0,2)
	g.Insert(0,3)
	g.Insert(1,4)
	g.Insert(1,5)
	g.Insert(3,6)

	assert.Equal(t, []int{0, 1, 4, 5, 2, 3, 6}, TopSortDFS(g))
}