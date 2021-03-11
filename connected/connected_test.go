package connected

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stevequadros/go-graph-theory/graph"
)

// find components in an undirected graph
// create graph with 3 components - 
// component 0
// - 0: 1,2
// - 1: 0,2
// - 2: 0,1
// component 1
// - 3: 4
// - 4: 3
// component 2
// - 5
func TestDFS(t *testing.T) {
	graph := graph.New(6, &graph.Options{Directed: false})
	graph.Insert(0,1)
	graph.Insert(0,2)
	graph.Insert(1,2)
	graph.Insert(1,0)
	graph.Insert(2,1)
	graph.Insert(3,4)

	res := Components(graph)
	assert.Equal(t, 3, res.Count)

	// should be in same component
	assert.Equal(t, 1, res.Components[0])
	assert.Equal(t, 1, res.Components[2])
	assert.Equal(t, 1, res.Components[1])

	assert.Equal(t, 2, res.Components[4])

	// node 5 is own component
	assert.Equal(t, 3, res.Components[5])

}