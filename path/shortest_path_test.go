package path

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

func TestSSSP(t *testing.T) {
	g := graph.New(8, &graph.Options{Directed: true, Weighted: true})
	g.InsertWeighted(0,2,3)
	g.InsertWeighted(0,1,6)
	g.InsertWeighted(1,4,8)
	g.InsertWeighted(1,5,11)
	g.InsertWeighted(2,1,4)
	g.InsertWeighted(2,4,4)
	g.InsertWeighted(2,3,11)
	g.InsertWeighted(3,6,9)
	g.InsertWeighted(4,3,-4)
	g.InsertWeighted(4,5,2)
	g.InsertWeighted(4,6,1)
	g.InsertWeighted(4,7,5)
	g.InsertWeighted(5,6,2)
	g.InsertWeighted(7,6,1)

	dist := DAGSP(g, 0) // 0,6,3,3,7,9,11,12
	assert.Equal(t, 7, dist[4])
	assert.Equal(t, 6, dist[1])
	assert.Equal(t, 12, dist[7])
}