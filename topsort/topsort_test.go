package topsort

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stevequadros/go-graph-theory/graph"
)

func TestTopSort(t *testing.T) {
	g := graph.New(6, &graph.Options{Directed: true})
	g.Insert(0,1)
	g.Insert(0,2)
	g.Insert(0,3)
	g.Insert(1,4)
	g.Insert(1,5)
	g.Insert(3,5)
	g.Insert(5,4)

	assert.Equal(t, []int{0,3,2,1,5,4}, TopsortKahn(g))

	t.Run("with cycle returns nil", func(t *testing.T) {
		// create cycle
		g.Insert(4,5)

		assert.Equal(t, []int(nil), TopsortKahn(g))
	})
}