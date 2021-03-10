package graph_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stevequadros/go-graph-theory/graph"
)

func TestNewGraph(t *testing.T) {
	g := graph.New(false)
	assert.NotNil(t, g)
}

func TestInsertNode(t *testing.T) {
	t.Run("directed", func(t *testing.T) {
		g := graph.New(true)

		g.Insert(1,2)
		assert.Len(t, g.Nodes, 2)

		g.Insert(2,1)

		node1 := g.Nodes[1]
		assert.Equal(t, 2, node1.Val)

		node2 := g.Nodes[2]
		assert.Equal(t, 1, node2.Val)
	})

	t.Run("undirected should insert an edge each way", func(t *testing.T) {
		g := graph.New(false)
		g.Insert(1,2)

		node1 := g.Nodes[1]
		assert.Equal(t, 2, node1.Val)

		node2 := g.Nodes[2]
		assert.Equal(t, 1, node2.Val)
	})	
}

func TestString(t *testing.T) {
	expected := `0 => 2 -> 1 -> 
1 => 3 -> 2 -> 
2 => 
3 => 
`

	g := graph.New(true)
	g.Insert(0,1)
	g.Insert(0,2)
	g.Insert(1,2)
	g.Insert(1,3)

	assert.Equal(t, expected, g.String())
}