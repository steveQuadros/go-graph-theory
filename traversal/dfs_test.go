package traversal

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
	graph := make([]*Node, 6)
	graph[0] = &Node{Val:1, Next: &Node{Val:2}}
	graph[1] = &Node{Val:0, Next: &Node{Val:2}}
	graph[2] = &Node{Val:0, Next: &Node{Val:1}}
	graph[3] = &Node{Val:4}
	graph[4] = &Node{Val:3}

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