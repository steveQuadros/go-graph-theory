package traversal

import (
	"testing"
	"github.com/stevequadros/go-graph-theory/graph"
	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	g := graph.New(11, &graph.Options{Directed: true})
	g.Insert(0,1)
	g.Insert(0,2)
	g.Insert(0,6)
	g.Insert(1,5)
	g.Insert(1,4)
	g.Insert(2,3)
	g.Insert(4,7)
	g.Insert(4,9)
	g.Insert(5,8)
	g.Insert(6,7)
	g.Insert(7,10)
	g.Insert(9,10)
	assert.Equal(t, []int{0,6,7,10}, ShortestPathBFS(g, 0, 10))
	assert.Equal(t, []int{}, ShortestPathBFS(g, 1, 6))
}

func TestShortestPathGrid(t *testing.T) {
	
}